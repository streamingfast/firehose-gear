package rpc

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/type/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FetchedBlockData struct {
	latestFinalizedHead uint64
	blockHash           types.Hash
	block               *types.SignedBlock
	events              []*parser.Event
}

type Fetcher struct {
	gearClients *firecoreRPC.Clients[*Client]

	fetchInterval            time.Duration
	latestBlockRetryInterval time.Duration
	logger                   *zap.Logger
	latestBlockNum           uint64
}

func NewFetcher(
	gearClients *firecoreRPC.Clients[*Client],
	fetchInterval time.Duration,
	latestBlockRetryInterval time.Duration,
	logger *zap.Logger) *Fetcher {
	return &Fetcher{
		gearClients:              gearClients,
		fetchInterval:            fetchInterval,
		latestBlockRetryInterval: latestBlockRetryInterval,
		logger:                   logger,
	}
}

func (f *Fetcher) IsBlockAvailable(blockNum uint64) bool {
	return blockNum <= f.latestBlockNum
}

func (f *Fetcher) Fetch(ctx context.Context, requestBlockNum uint64) (b *pbbstream.Block, skipped bool, err error) {
	f.logger.Info("fetching block", zap.Uint64("block_num", requestBlockNum))

	sleepDuration := time.Duration(0)
	//TODO: move this logic in the firecore binary, as we do this in all the fetchers
	for f.latestBlockNum < requestBlockNum {
		time.Sleep(sleepDuration)

		f.latestBlockNum, err = f.fetchLatestBlockNum(ctx)
		if err != nil {
			return nil, false, fmt.Errorf("fetching latest block num: %w", err)
		}

		f.logger.Info("got latest block num", zap.Uint64("latest_block_num", f.latestBlockNum), zap.Uint64("requested_block_num", requestBlockNum))

		if f.latestBlockNum >= requestBlockNum {
			break
		}
		sleepDuration = f.latestBlockRetryInterval
	}

	blockData, err := f.fetchBlockData(ctx, requestBlockNum)
	if err != nil {
		return nil, false, fmt.Errorf("fetching block data: %w", err)
	}

	f.logger.Info("converting block", zap.Uint64("block_num", requestBlockNum))
	bstreamBlock, err := convertBlock(blockData)
	if err != nil {
		return nil, false, fmt.Errorf("converting block %d from rpc response: %w", requestBlockNum, err)
	}

	return bstreamBlock, false, nil
}

func (f *Fetcher) fetchBlockData(_ context.Context, requestedBlockNum uint64) (*FetchedBlockData, error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) (*FetchedBlockData, error) {
		latestFinalizedHead, err := client.client.RPC.Chain.GetFinalizedHead()
		if err != nil {
			return nil, fmt.Errorf("unable to fetch latest finalized head: %w", err)
		}

		headBlock, err := client.client.RPC.Chain.GetBlock(latestFinalizedHead)
		if err != nil {
			return nil, fmt.Errorf("failed to get block: %w", err)
		}

		blockHash, err := client.client.RPC.Chain.GetBlockHash(requestedBlockNum)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch latest block header: %w", err)
		}

		block, err := client.client.RPC.Chain.GetBlock(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get block: %w", err)
		}

		fetchedBlockData := &FetchedBlockData{
			latestFinalizedHead: uint64(headBlock.Block.Header.Number),
			blockHash:           blockHash,
			block:               block,
		}

		if requestedBlockNum > 0 {
			events, err := client.eventsRetriever.GetEvents(blockHash)
			if err != nil {
				return nil, fmt.Errorf("failed to get events: %w", err)
			}
			fetchedBlockData.events = events
		}

		return fetchedBlockData, nil
	})
}

func (f *Fetcher) fetchLatestBlockNum(_ context.Context) (uint64, error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) (uint64, error) {
		header, err := client.client.RPC.Chain.GetHeaderLatest()
		if err != nil {
			return 0, fmt.Errorf("unable to fetch latest block header: %w", err)
		}
		return uint64(header.Number), nil
	})
}

func convertBlock(data *FetchedBlockData) (*pbbstream.Block, error) {
	convertedEvents, err := convertEvents(data.events)
	if err != nil {
		return nil, fmt.Errorf("failed to convert events: %w", err)
	}

	b := data.block
	latestFinalizedHead := data.latestFinalizedHead
	blockHash := data.blockHash

	block := &pbgear.Block{
		Number: uint64(b.Block.Header.Number),
		Hash:   blockHash.Hex(),
	}

	timestamp := big.NewInt(0)

	if b.Block.Header.Number > 0 {
		block.Header = convertHeader(b)
		block.Extrinsics = convertExtrinsics(b.Block.Extrinsics)
		block.Events = convertedEvents
		block.DigestItems = convertLogs(b.Block.Header.Digest)
		block.Justification = b.Justification

		timestampBytes := b.Block.Extrinsics[0].Method.Args
		timestampDecoder := scale.NewDecoder(bytes.NewBuffer(timestampBytes))
		timestamp, err = timestampDecoder.DecodeUintCompact()
		if err != nil {
			return nil, fmt.Errorf("unable to decode timestamp: %w", err)
		}
	}

	anyBlock, err := anypb.New(block)
	if err != nil {
		return nil, fmt.Errorf("unable to create anypb: %w", err)
	}

	var parentBlockNum uint64
	if block.Number > 0 {
		parentBlockNum = block.Number - 1
	}

	libNum := uint64(0)
	if latestFinalizedHead > block.Number && block.Number > 0 {
		libNum = block.Number - 1
	}

	parentHash := ""
	if block.Number > 0 {
		parentHash = block.Header.ParentHash
	}

	if block.Number == 0 {
		parentHash = "0x0000000000000000000000000000000000000000000000000000000000000000"
	}

	bstreamBlock := &pbbstream.Block{
		Number:    block.Number,
		Id:        block.Hash,
		ParentId:  parentHash,
		Timestamp: timestamppb.New(time.Unix(timestamp.Int64(), 0)),
		LibNum:    libNum,
		ParentNum: parentBlockNum,
		Payload:   anyBlock,
	}

	return bstreamBlock, nil
}

func convertHeader(b *types.SignedBlock) *pbgear.Header {
	return &pbgear.Header{
		ParentHash:     b.Block.Header.ParentHash.Hex(),
		StateRoot:      b.Block.Header.StateRoot.Hex(),
		ExtrinsicsRoot: b.Block.Header.ExtrinsicsRoot.Hex(),
	}
}

func convertLogs(digestItems []types.DigestItem) []*pbgear.DigestItem {
	digest := make([]*pbgear.DigestItem, len(digestItems))
	for i, item := range digestItems {

		digestItem := &pbgear.DigestItem{}

		if item.IsChangesTrieRoot {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsChangesTrieRoot{
					AsChangesTrieRoot: item.AsChangesTrieRoot.Hex(),
				},
			}
		}

		if item.IsPreRuntime {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsPreRuntime{
					AsPreRuntime: convertPreRuntime(item.AsPreRuntime),
				},
			}
		}

		if item.IsConsensus {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsConsensus{
					AsConsensus: convertAsConsensus(item.AsConsensus),
				},
			}
		}

		if item.IsSeal {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsSeal{
					AsSeal: convertSeal(item.AsSeal),
				},
			}
		}

		if item.IsChangesTrieSignal {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsChangesTrieSignal{
					AsChangesTrieSignal: convertChangesTrieSignal(item.AsChangesTrieSignal),
				},
			}
		}

		if item.IsOther {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsOther{
					AsOther: item.AsOther,
				},
			}
		}

		digest[i] = digestItem
	}
	return digest
}

func convertPreRuntime(preRuntime types.PreRuntime) *pbgear.PreRuntime {
	return &pbgear.PreRuntime{
		ConsensusEngineId: uint32(preRuntime.ConsensusEngineID),
		Bytes:             preRuntime.Bytes,
	}
}

func convertAsConsensus(consensus types.Consensus) *pbgear.Consensus {
	return &pbgear.Consensus{
		ConsensusEngineId: uint32(consensus.ConsensusEngineID),
		Bytes:             consensus.Bytes,
	}

}

func convertSeal(seal types.Seal) *pbgear.Seal {
	return &pbgear.Seal{
		ConsensusEngineId: uint32(seal.ConsensusEngineID),
		Bytes:             seal.Bytes,
	}
}

func convertChangesTrieSignal(signal types.ChangesTrieSignal) *pbgear.ChangesTrieSignal {
	return &pbgear.ChangesTrieSignal{
		IsNewConfiguration: signal.IsNewConfiguration,
		AsNewConfiguration: signal.AsNewConfiguration,
	}
}

func convertExtrinsics(extrinsics []types.Extrinsic) []*pbgear.Extrinsic {
	gearExtrinsics := make([]*pbgear.Extrinsic, 0, len(extrinsics))
	for _, extrinsic := range extrinsics {
		gearExtrinsics = append(gearExtrinsics, &pbgear.Extrinsic{
			Version:   uint32(extrinsic.Version),
			Signature: convertExtrinsicSignature(extrinsic.Signature),
			Method:    convertExtrinsicCall(extrinsic.Method),
		})
	}
	return gearExtrinsics
}

func convertExtrinsicSignature(signature types.ExtrinsicSignatureV4) *pbgear.Signature {
	nonce := big.Int(signature.Nonce)
	tip := big.Int(signature.Tip)
	return &pbgear.Signature{
		Signer:    convertMultiAddress(signature.Signer),
		Signature: convertMultiSignature(signature.Signature),
		Era:       convertExtrinsicEra(signature.Era),
		Nonce:     nonce.String(),
		Tip:       tip.String(),
	}
}

func convertExtrinsicCall(call types.Call) *pbgear.Call {
	return &pbgear.Call{
		CallIndex: convertExtrinsicCallIndex(call.CallIndex),
		Args:      call.Args,
	}
}

func convertExtrinsicCallIndex(callIndex types.CallIndex) *pbgear.CallIndex {
	return &pbgear.CallIndex{
		SectionIndex: uint32(callIndex.SectionIndex),
		MethodIndex:  uint32(callIndex.MethodIndex),
	}
}

func convertMultiSignature(signature types.MultiSignature) *pbgear.MultiSignature {
	return &pbgear.MultiSignature{
		IsEd_25519: signature.IsEcdsa,
		AsEd_25519: signature.AsEd25519.Hex(),
		IsSr_25519: signature.IsSr25519,
		AsSr_25519: signature.AsSr25519.Hex(),
		IsEcdsa:    signature.IsEcdsa,
		AsEcdsa:    signature.AsEcdsa.Hex(),
	}
}

func convertMultiAddress(signer types.MultiAddress) *pbgear.MultiAddress {
	return &pbgear.MultiAddress{
		IsId:         signer.IsID,
		AsId:         signer.AsID.ToHexString(),
		IsIndex:      signer.IsIndex,
		AsIndex:      uint32(signer.AsIndex),
		IsRaw:        signer.IsRaw,
		AsRaw:        string(signer.AsRaw),
		IsAddress_32: signer.IsAddress32,
		AsAddress_32: string(signer.AsAddress32[:]),
		IsAddress_20: signer.IsAddress20,
		AsAddress_20: string(signer.AsAddress20[:]),
	}
}

func convertExtrinsicEra(era types.ExtrinsicEra) *pbgear.ExtrinsicEra {
	if &era == nil {
		return nil
	}
	return &pbgear.ExtrinsicEra{
		IsImmortalEra: era.IsImmortalEra,
		IsMortalEra:   era.IsMortalEra,
		AsMortalEra:   convertMortalEra(era.AsMortalEra),
	}
}

func convertMortalEra(era types.MortalEra) *pbgear.MortalEra {
	return &pbgear.MortalEra{
		First:  uint32(era.First),
		Second: uint32(era.Second),
	}
}

func convertNone(nonce types.UCompact) string {
	b := big.Int(nonce)
	return b.String()
}

func convertPaymentFields(paymentFields generic.DefaultPaymentFields) *pbgear.PaymentFields {
	b := big.Int(paymentFields.Tip)
	return &pbgear.PaymentFields{
		Tip: b.String(),
	}
}

func convertEvents(events []*parser.Event) ([]*pbgear.Event, error) {
	pbgearEvent := make([]*pbgear.Event, 0, len(events))
	for _, evt := range events {
		// TODO: add the fields as bytes directly, and decode them in the substreams
		// fields, err := convertDecodedFields(evt.Fields)
		// if err != nil {
		// 	return nil, err
		// }
		pbgearEvent = append(pbgearEvent, &pbgear.Event{
			Name: evt.Name,
			// Fields: fields,
			Id:     string(evt.EventID[:]),
			Phase:  convertPhase(evt.Phase),
			Topics: convertTopics(evt.Topics),
		})
	}
	return pbgearEvent, nil
}

func convertPhase(phase *types.Phase) *pbgear.Phase {
	return &pbgear.Phase{
		IsApplyExtrinsic: phase.IsApplyExtrinsic,
		AsApplyExtrinsic: phase.AsApplyExtrinsic,
		IsFinalization:   phase.IsFinalization,
		IsInitialization: phase.IsInitialization,
	}

}
func convertTopics(hashes []types.Hash) []string {
	topics := make([]string, 0, len(hashes))
	for _, hash := range hashes {
		topics = append(topics, hash.Hex())
	}
	return topics
}
