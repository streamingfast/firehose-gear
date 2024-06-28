package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	retriever "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/type/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Fetcher struct {
	gearClients        *firecoreRPC.Clients[*Client]
	extrinsicRetriever *firecoreRPC.Clients[retriever.DefaultExtrinsicRetriever]
	eventsRetriever    *firecoreRPC.Clients[retriever.EventRetriever]

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

	extrinsics, err := f.fetchExtrinsics(ctx, blockHash)
	if err != nil {
		return nil, false, fmt.Errorf("fetching extrinsics for block %d: %w", requestBlockNum, err)
	}

	events, err := f.fetchEvents(ctx, blockHash)
	if err != nil {
		return nil, false, fmt.Errorf("fetching events for block %d: %w", requestBlockNum, err)
	}

	f.logger.Info("converting block", zap.Uint64("block_num", requestBlockNum))
	bstreamBlock, err := convertBlock(blockHash, block, extrinsics, events)
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
	return firecoreRPC.WithClients(f.extrinsicRetriever, func(client retriever.DefaultExtrinsicRetriever) ([]*parser.Extrinsic[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields], error) {
		extrinsincs, err := client.GetExtrinsics(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get extrinsics: %w", err)
		}
		return extrinsincs, nil
	})
}

func (f *Fetcher) fetchEvents(_ context.Context, blockHash types.Hash) ([]*parser.Event, error) {
	return firecoreRPC.WithClients(f.eventsRetriever, func(client retriever.EventRetriever) ([]*parser.Event, error) {
		events, err := client.GetEvents(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get events: %w", err)
		}
		return events, nil
	})
}

func convertBlock(
	blockHash types.Hash,
	b *types.SignedBlock,
	extrinsics []*parser.Extrinsic[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields],
	events []*parser.Event,
) (*pbbstream.Block, error) {
	block := &pbgear.Block{
		Number:        uint64(b.Block.Header.Number),
		Hash:          blockHash[:],
		Header:        convertHeader(b),
		Extrinsics:    convertExtrinsics(extrinsics),
		Events:        convertEvents(events),
		Justification: b.Justification,
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

	bstreamBlock := &pbbstream.Block{
		Number:   block.Number,
		Id:       string(block.Hash),
		ParentId: string(block.Header.ParentHash),
		// the timestamp needs to be fetched from the extrinsics and get the item called timestamp (set)
		Timestamp: timestamppb.New(time.Unix(int64(fetchTimestampFromExtrinsics(extrinsics)), 0)),
		LibNum:    libNum,
		ParentNum: parentBlockNum,
		Payload:   anyBlock,
	}

	return bstreamBlock, nil
}

const TIMESTAMP_SET_NAME = "Timestamp.set"
const TIMESTAMP_FIELD_NAME = "now"

func fetchTimestampFromExtrinsics(extrinsics []*parser.Extrinsic[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields]) uint64 {
	for _, extrinsic := range extrinsics {
		if extrinsic.Name == TIMESTAMP_SET_NAME {
			for _, callField := range extrinsic.CallFields {
				if callField.Name == TIMESTAMP_FIELD_NAME {
					return callField.Value.(uint64)
				}
			}
		}
	}
	panic("timestamp not found in extrinsics")
}

func convertHeader(b *types.SignedBlock) *pbgear.Header {
	return &pbgear.Header{
		ParentHash:     b.Block.Header.ParentHash[:],
		StateRoot:      b.Block.Header.StateRoot[:],
		ExtrinsicsRoot: b.Block.Header.ExtrinsicsRoot[:],
		Digest:         convertDigest(b.Block.Header.Digest),
	}
}

func convertDigest(digestItems []types.DigestItem) []*pbgear.DigestItem {
	digest := make([]*pbgear.DigestItem, len(digestItems))
	for i, item := range digestItems {
		digest[i] = &pbgear.DigestItem{
			IsChangesTrieRoot:   item.IsChangesTrieRoot,
			AsChangesTrieRoot:   item.AsChangesTrieRoot[:],
			IsPreRuntime:        item.IsPreRuntime,
			AsPreRuntime:        convertPreRuntime(item.AsPreRuntime),
			IsConsensus:         item.IsConsensus,
			AsConsensus:         convertAsConsensus(item.AsConsensus),
			IsSeal:              item.IsSeal,
			AsSeal:              convertSeal(item.AsSeal),
			IsChangesTrieSignal: item.IsChangesTrieSignal,
			AsChangesTrieSignal: convertChangesTrieSignal(item.AsChangesTrieSignal),
			IsOther:             item.IsOther,
			AsOther:             item.AsOther,
		}
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

func convertExtrinsics(extrinsics []*parser.Extrinsic[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields]) []*pbgear.Extrinsic {
	gearExtrinsics := make([]*pbgear.Extrinsic, 0, len(extrinsics))
	for _, extrinsic := range extrinsics {
		gearExtrinsics = append(gearExtrinsics, &pbgear.Extrinsic{
			Name:       extrinsic.Name,
			CallFields: convertDecodedFields(extrinsic.CallFields),
			CallIndex:  convertCallIndex(extrinsic.CallIndex),
			Version:    uint32(extrinsic.Version),
			Signature:  convertExtrinsicsSignature(extrinsic.Signature),
		})
	}

	return gearExtrinsics
}

func convertDecodedFields(callFields []*registry.DecodedField) []*pbgear.Field {
	fields := make([]*pbgear.Field, 0, len(callFields))
	for _, callField := range callFields {
		fields = append(fields, &pbgear.Field{
			Name:        callField.Name,
			Value:       callField.Value.(*anypb.Any),
			LookupIndex: callField.LookupIndex,
		})
	}
	return fields
}

func convertCallIndex(callIndex types.CallIndex) *pbgear.CallIndex {
	return &pbgear.CallIndex{
		SectionIndex: uint32(callIndex.SectionIndex),
		MethodIndex:  uint32(callIndex.MethodIndex),
	}
}

func convertExtrinsicsSignature(signature generic.GenericExtrinsicSignature[types.MultiAddress, types.MultiSignature, generic.DefaultPaymentFields]) *pbgear.Signature {
	pbgearSignature := &pbgear.Signature{}

	// todo: convert the signatures

	return pbgearSignature
}

func convertEvents(events []*parser.Event) []*pbgear.Event {
	pbgearEvent := make([]*pbgear.Event, 0, len(events))
	for _, evt := range events {
		pbgearEvent = append(pbgearEvent, &pbgear.Event{
			Name:   evt.Name,
			Fields: convertDecodedFields(evt.Fields),
			Id:     string(evt.EventID[:]),
			Phase:  convertPhase(evt.Phase),
			Topics: convertTopics(evt.Topics),
		})
	}
	return pbgearEvent
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
