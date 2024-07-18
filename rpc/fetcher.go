package rpc

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	cclient "github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
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
	metadata            []byte
}

type LastBlockInfo struct {
	blockNum    uint64
	blockHash   types.Hash
	specVersion uint32
}

type Fetcher struct {
	gearClients *firecoreRPC.Clients[*Client]

	fetchInterval            time.Duration
	latestBlockRetryInterval time.Duration
	logger                   *zap.Logger

	latestBlockNum uint64

	lastBlockInfo *LastBlockInfo
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
		lastBlockInfo:            &LastBlockInfo{},
	}
}

func (f *Fetcher) IsBlockAvailable(blockNum uint64) bool {
	return blockNum <= f.latestBlockNum
}

func (f *Fetcher) Fetch(ctx context.Context, requestBlockNum uint64) (b *pbbstream.Block, skipped bool, err error) {
	f.logger.Info("gear fetching block", zap.Uint64("block_num", requestBlockNum))

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
		f.logger.Warn("fetching block data", zap.Uint64("block_num", requestBlockNum), zap.Error(err))
		return nil, false, fmt.Errorf("fetching block data: %w", err)
	}

	f.logger.Info("converting block", zap.Uint64("block_num", requestBlockNum), zap.Uint32("spec_version", f.lastBlockInfo.specVersion))
	bstreamBlock, err := convertBlock(blockData, f.lastBlockInfo.specVersion)
	if err != nil {
		f.logger.Warn("converting block", zap.Uint64("block_num", requestBlockNum), zap.Error(err))
		return nil, false, fmt.Errorf("converting block %d from rpc response: %w", requestBlockNum, err)
	}

	return bstreamBlock, false, nil
}

func (f *Fetcher) fetchBlockData(_ context.Context, requestedBlockNum uint64) (*FetchedBlockData, error) {
	return firecoreRPC.WithClients(f.gearClients, func(client *Client) (*FetchedBlockData, error) {
		f.logger.Info("fetching block data", zap.Uint64("block_num", requestedBlockNum))
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

		runtimeVersion, err := client.state.GetRuntimeVersion(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get runtime version at block hash %s: %w", blockHash.Hex(), err)
		}

		requestedBlockSpecVersion := uint32(runtimeVersion.SpecVersion)
		parentSpecVersion := uint32(0)

		shouldFetchParentVersion := true
		if block.Block.Header.ParentHash == f.lastBlockInfo.blockHash {
			shouldFetchParentVersion = false
			parentSpecVersion = f.lastBlockInfo.specVersion
		}

		if requestedBlockNum == 0 {
			shouldFetchParentVersion = false
		}

		f.logger.Info(
			"requested block spec version",
			zap.Uint32("spec_version", requestedBlockSpecVersion),
			zap.Uint64("block_num", requestedBlockNum),
		)

		if shouldFetchParentVersion {
			pSpecVersion, err := f.fetchParentSpecVersion(client, block.Block.Header.ParentHash)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch parent spec version: %w", err)
			}
			f.logger.Info("fetched parent spec version", zap.Uint32("spec_version", pSpecVersion), zap.Uint64("block_num", requestedBlockNum-1))
			parentSpecVersion = pSpecVersion
		}

		if shouldUpdateMetadata(requestedBlockSpecVersion, parentSpecVersion, isForward(f.lastBlockInfo.blockNum, requestedBlockNum)) {
			err := cclient.CallWithBlockHash(client.client.Client, &fetchedBlockData.metadata, "state_getMetadata", &blockHash)
			if err != nil {
				return nil, fmt.Errorf("failed to get metadata: %w", err)
			}

			f.logger.Info("metadata fetched", zap.Uint32("spec_version", requestedBlockSpecVersion), zap.Uint64("block_num", requestedBlockNum))
		}

		f.lastBlockInfo = &LastBlockInfo{
			blockNum:    requestedBlockNum,
			blockHash:   blockHash,
			specVersion: requestedBlockSpecVersion,
		}

		return fetchedBlockData, nil
	})
}

func (f *Fetcher) fetchParentSpecVersion(client *Client, parentBlockHash types.Hash) (uint32, error) {
	parentRuntimeVersion, err := client.state.GetRuntimeVersion(parentBlockHash)
	if err != nil {
		return 0, fmt.Errorf("failed to get runtime version at block hash %s: %w", parentBlockHash.Hex(), err)
	}

	return uint32(parentRuntimeVersion.SpecVersion), nil
}

func isForward(lastProcessedBlockNum, requestedBlockNum uint64) bool {
	return lastProcessedBlockNum <= requestedBlockNum
}

func shouldUpdateMetadata(specVersion uint32, parentSpecVersion uint32, isForward bool) bool {
	if !isForward {
		return false
	}

	return specVersion != parentSpecVersion
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

func convertBlock(data *FetchedBlockData, specVersion uint32) (*pbbstream.Block, error) {
	convertedEvents, err := convertEvents(data.events)
	if err != nil {
		return nil, fmt.Errorf("failed to convert events: %w", err)
	}

	b := data.block
	latestFinalizedHead := data.latestFinalizedHead
	blockHash := data.blockHash

	block := &pbgear.Block{
		Number:        uint64(b.Block.Header.Number),
		Hash:          blockHash[:],
		Header:        convertHeader(data, specVersion, data.metadata),
		Extrinsics:    convertExtrinsics(b.Block.Extrinsics),
		Events:        convertedEvents,
		DigestItems:   convertLogs(b.Block.Header.Digest),
		Justification: b.Justification,
	}

	timestamp := big.NewInt(0)
	if b.Block.Header.Number > 0 {
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

	bstreamBlock := &pbbstream.Block{
		Number:    block.Number,
		Id:        hex.EncodeToString(block.Hash),
		ParentId:  hex.EncodeToString(block.Header.ParentHash),
		Timestamp: timestamppb.New(time.UnixMilli(timestamp.Int64())),
		LibNum:    libNum,
		ParentNum: parentBlockNum,
		Payload:   anyBlock,
	}

	return bstreamBlock, nil
}

func convertHeader(fetchedBlockData *FetchedBlockData, specVersion uint32, metadata []byte) *pbgear.Header {
	h := &pbgear.Header{
		ParentHash:     fetchedBlockData.block.Block.Header.ParentHash[:],
		StateRoot:      fetchedBlockData.block.Block.Header.StateRoot[:],
		ExtrinsicsRoot: fetchedBlockData.block.Block.Header.ExtrinsicsRoot[:],
		SpecVersion:    specVersion,
	}

	// set the metadata if the previous spec version is different from the current spec version
	// this means that the metadata has been updated
	if len(metadata) > 0 {
		h.UpdatedMetadata = metadata
	}

	return h
}

func convertLogs(digestItems []types.DigestItem) []*pbgear.DigestItem {
	digest := make([]*pbgear.DigestItem, len(digestItems))
	for i, item := range digestItems {

		digestItem := &pbgear.DigestItem{}

		if item.IsChangesTrieRoot {
			digestItem = &pbgear.DigestItem{
				Item: &pbgear.DigestItem_AsChangesTrieRoot{
					AsChangesTrieRoot: item.AsChangesTrieRoot[:],
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
		AsEd_25519: signature.AsEd25519[:],
		IsSr_25519: signature.IsSr25519,
		AsSr_25519: signature.AsSr25519[:],
		IsEcdsa:    signature.IsEcdsa,
		AsEcdsa:    signature.AsEcdsa[:],
	}
}

func convertMultiAddress(signer types.MultiAddress) *pbgear.MultiAddress {
	return &pbgear.MultiAddress{
		IsId:         signer.IsID,
		AsId:         signer.AsID[:],
		IsIndex:      signer.IsIndex,
		AsIndex:      uint32(signer.AsIndex),
		IsRaw:        signer.IsRaw,
		AsRaw:        signer.AsRaw,
		IsAddress_32: signer.IsAddress32,
		AsAddress_32: signer.AsAddress32[:],
		IsAddress_20: signer.IsAddress20,
		AsAddress_20: signer.AsAddress20[:],
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
	for i, evt := range events {
		fmt.Println("EVENT NUM", i)
		fields, err := convertEventFields(evt.Fields)
		if err != nil {
			return nil, err
		}

		pbgearEvent = append(pbgearEvent, &pbgear.Event{
			Name:   evt.Name,
			Fields: fields,
			Id:     evt.EventID[:],
			Phase:  convertPhase(evt.Phase),
			Topics: convertTopics(evt.Topics),
		})
	}
	return pbgearEvent, nil
}

func convertEventFields(fields registry.DecodedFields) ([][]byte, error) {
	out := make([][]byte, 0)
	for i, field := range fields {
		buffer := bytes.NewBuffer(nil)
		fieldEncoder := scale.NewEncoder(buffer)
		err := fieldEncoder.Encode(field.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to encode field: %w", err)
		}
		out = append(out, buffer.Bytes())
		fmt.Printf("event %d field: %s type: %s\n", i, hex.EncodeToString(buffer.Bytes()), field.Name)
		b, err := json.MarshalIndent(field.Value, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal field: %w", err)
		}
		fmt.Println(string(b))
	}
	return out, nil
}

func convertPhase(phase *types.Phase) *pbgear.Phase {
	return &pbgear.Phase{
		IsApplyExtrinsic: phase.IsApplyExtrinsic,
		AsApplyExtrinsic: phase.AsApplyExtrinsic,
		IsFinalization:   phase.IsFinalization,
		IsInitialization: phase.IsInitialization,
	}

}
func convertTopics(hashes []types.Hash) [][]byte {
	topics := make([][]byte, 0, len(hashes))
	for _, hash := range hashes {
		topics = append(topics, hash[:])
	}
	return topics
}
