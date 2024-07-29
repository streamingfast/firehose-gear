package rpc

import (
	"bytes"
	"context"
	"log"
	"testing"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	pbgear "github.com/streamingfast/firehose-gear/pb/sf/gear/type/v1"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func Test_DecodeBlockExtrinsics(t *testing.T) {
	latestBlockRetryInterval := time.Second
	rpcEndpoint := "https://vara-mainnet.public.blastapi.io"
	api, err := gsrpc.NewSubstrateAPI(rpcEndpoint)
	require.NoError(t, err)

	gearClients := firecoreRPC.NewClients[*Client]()
	fetchInterval := time.Second
	logger := zap.NewNop()

	gearClient, err := NewClient(rpcEndpoint)
	require.NoError(t, err)
	gearClients.Add(gearClient)

	rpcFetcher := NewFetcher(gearClients, fetchInterval, latestBlockRetryInterval, logger)
	b, _, err := rpcFetcher.Fetch(context.Background(), 14090409)
	require.NoError(t, err)

	pbGearBlock := &pbgear.Block{}
	err = proto.Unmarshal(b.Payload.Value, pbGearBlock)
	require.NoError(t, err)
	require.NotNil(t, pbGearBlock)

	metadata, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}

	factory := registry.NewFactory()
	callRegistry, err := factory.CreateCallRegistry(metadata)
	require.NoError(t, err)

	for _, extrinsic := range pbGearBlock.Extrinsics {
		decodedFields := decodeCallExtrinsics(t, callRegistry, extrinsic)
		_ = decodedFields
	}
}

func Test_DecodeBlockEvents(t *testing.T) {
	latestBlockRetryInterval := time.Second
	rpcEndpoint := "https://vara-mainnet.public.blastapi.io"
	api, err := gsrpc.NewSubstrateAPI(rpcEndpoint)
	require.NoError(t, err)

	gearClients := firecoreRPC.NewClients[*Client]()
	fetchInterval := time.Second
	logger := zap.NewNop()

	gearClient, err := NewClient(rpcEndpoint)
	require.NoError(t, err)
	gearClients.Add(gearClient)

	rpcFetcher := NewFetcher(gearClients, fetchInterval, latestBlockRetryInterval, logger)
	b, _, err := rpcFetcher.Fetch(context.Background(), 14067000)
	require.NoError(t, err)

	pbGearBlock := &pbgear.Block{}
	err = proto.Unmarshal(b.Payload.Value, pbGearBlock)
	require.NoError(t, err)
	require.NotNil(t, pbGearBlock)

	metadata, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}

	factory := registry.NewFactory()
	eventRegistry, err := factory.CreateEventRegistry(metadata)
	require.NoError(t, err)

	decodedFields := decodeEvents(t, eventRegistry, pbGearBlock.RawEvents)
	_ = decodedFields
}

func convertCallIndex(ci *pbgear.CallIndex) types.CallIndex {
	return types.CallIndex{
		SectionIndex: uint8(ci.SectionIndex),
		MethodIndex:  uint8(ci.MethodIndex),
	}
}

func decodeCallExtrinsics(t *testing.T, callRegistry registry.CallRegistry, extrinsic *pbgear.Extrinsic) registry.DecodedFields {
	callIndex := extrinsic.Method.CallIndex
	args := extrinsic.Method.Args

	callDecoder, ok := callRegistry[convertCallIndex(callIndex)]
	require.True(t, ok)

	decoder := scale.NewDecoder(bytes.NewReader(args))

	callFields, err := callDecoder.Decode(decoder)
	require.NoError(t, err)
	return callFields
}

func decodeEvents(t *testing.T, eventRegistry registry.EventRegistry, storageEvents []byte) []*parser.Event {
	decoder := scale.NewDecoder(bytes.NewReader(storageEvents))

	eventsCount, err := decoder.DecodeUintCompact()
	require.NoError(t, err)

	var events []*parser.Event

	for i := uint64(0); i < eventsCount.Uint64(); i++ {
		var phase types.Phase

		err := decoder.Decode(&phase)
		require.NoError(t, err)

		var eventID types.EventID

		err = decoder.Decode(&eventID)
		require.NoError(t, err)

		eventDecoder, ok := eventRegistry[eventID]
		require.True(t, ok)

		eventFields, err := eventDecoder.Decode(decoder)
		require.NoError(t, err)

		var topics []types.Hash

		err = decoder.Decode(&topics)
		require.NoError(t, err)

		event := &parser.Event{
			Name:    eventDecoder.Name,
			Fields:  eventFields,
			EventID: eventID,
			Phase:   &phase,
			Topics:  topics,
		}

		events = append(events, event)
	}

	return events
}
