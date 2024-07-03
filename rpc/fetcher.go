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

	//TODO: how to compute block hash from block
	// or maybe should we compute it in the substreams foundational module?
	blockHash, err := f.fetchBlockHash(ctx, requestBlockNum)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get block hash for block %d: %w", requestBlockNum, err)
	}

	f.logger.Info("fetching block", zap.Uint64("block_num", requestBlockNum))
	block, err := f.fetchBlock(ctx, blockHash)
	if err != nil {
		return nil, false, fmt.Errorf("fetching block %d: %w", requestBlockNum, err)
	}
	f.logger.Debug("block fetched successfully", zap.Uint64("block_num", requestBlockNum))

	events, err := f.fetchEvents(ctx, blockHash)
	if err != nil {
		return nil, false, fmt.Errorf("fetching events for block %d: %w", requestBlockNum, err)
	}

	f.logger.Info("converting block", zap.Uint64("block_num", requestBlockNum))
	bstreamBlock, err := convertBlock(blockHash, block, events)
	if err != nil {
		return nil, false, fmt.Errorf("converting block %d from rpc response: %w", requestBlockNum, err)
	}

	return bstreamBlock, false, nil

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

func (f *Fetcher) fetchBlockHash(_ context.Context, requestedBlockNum uint64) (types.Hash, error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) (types.Hash, error) {
		var hash types.Hash
		hash, err := client.client.RPC.Chain.GetBlockHash(requestedBlockNum)
		if err != nil {
			return hash, fmt.Errorf("unable to fetch latest block header: %w", err)
		}
		return hash, nil
	})
}

func (f *Fetcher) fetchBlock(_ context.Context, blockHash types.Hash) (*types.SignedBlock, error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) (*types.SignedBlock, error) {
		block, err := client.client.RPC.Chain.GetBlock(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get block: %w", err)
		}
		return block, nil
	})
}

func (f *Fetcher) fetchExtrinsics(_ context.Context, blockHash types.Hash) ([]*parser.Extrinsic[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields], error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) ([]*parser.Extrinsic[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields], error) {
		extrinsincs, err := client.extrinsicsRetriever.GetExtrinsics(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get extrinsics: %w", err)
		}
		return extrinsincs, nil
	})
}

func (f *Fetcher) fetchEvents(_ context.Context, blockHash types.Hash) ([]*parser.Event, error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) ([]*parser.Event, error) {
		events, err := client.eventsRetriever.GetEvents(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get events: %w", err)
		}
		return events, nil
	})
}

func convertBlock(
	blockHash types.Hash,
	b *types.SignedBlock,
	events []*parser.Event,
) (*pbbstream.Block, error) {
	convertedEvents, err := convertEvents(events)
	if err != nil {
		return nil, fmt.Errorf("failed to convert events: %w", err)
	}

	block := &pbgear.Block{
		Number:        uint64(b.Block.Header.Number),
		Hash:          blockHash.Hex(),
		Header:        convertHeader(b),
		Extrinsics:    convertExtrinsics(b.Block.Extrinsics),
		Events:        convertedEvents,
		DigestItems:   convertLogs(b.Block.Header.Digest),
		Justification: b.Justification,
	}

	timestampBytes := b.Block.Extrinsics[0].Method.Args
	timestampDecoder := scale.NewDecoder(bytes.NewBuffer(timestampBytes))
	timestamp, err := timestampDecoder.DecodeUintCompact()
	if err != nil {
		return nil, fmt.Errorf("unable to decode timestamp: %w", err)
	}

	fmt.Println("timestamp: ", timestamp)

	anyBlock, err := anypb.New(block)
	if err != nil {
		return nil, fmt.Errorf("unable to create anypb: %w", err)
	}

	var parentBlockNum uint64
	if block.Number > 0 {
		parentBlockNum = block.Number - 1
	}

	libNum := uint64(0)

	bstreamBlock := &pbbstream.Block{
		Number:   block.Number,
		Id:       string(block.Hash),
		ParentId: string(block.Header.ParentHash),
		// the timestamp needs to be fetched from the extrinsics and get the item called timestamp (set)
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
