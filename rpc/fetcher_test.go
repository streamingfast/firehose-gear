package rpc

import (
	"context"
	"log"
	"testing"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
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
	_ = callRegistry

	decodedExtrinsics, err := extrinsicsToProto(callRegistry, pbGearBlock.Extrinsics)
	require.NoError(t, err)
	_ = decodedExtrinsics
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

	decodedFields, err := decodeEvents(eventRegistry, pbGearBlock.RawEvents)
	require.NoError(t, err)
	_ = decodedFields
}
