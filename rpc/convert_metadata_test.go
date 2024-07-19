package rpc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/types"
	"github.com/stretchr/testify/require"
)

func Test_ConvertMetadata(t *testing.T) {
	gearClients := firecoreRPC.NewClients[*Client]()
	gearClient := NewClient("https://vara-mainnet.public.blastapi.io")
	gearClients.Add(gearClient)

	mc := NewMetadataConverter(gearClients, nil, nil)
	ttypes, err := mc.Convert("")
	require.NoError(t, err)

	protobufMessages := &types.ProtobufMessages{
		SyntaxName:  "proto3",
		PackageName: "sf.gear.extrinsics.type.v1",
	}

	for key := range ttypes {
		protobufMessages.Messages = append(protobufMessages.Messages, &types.ProtobufMessage{
			Name: key,
		})
	}

	fmt.Println(protobufMessages.ToProtoMessage())
	require.True(t, false)
}

func Test_DecodeField(t *testing.T) {
	url := "https://vara-mainnet.public.blastapi.io" // Replace with the actual URL of your Gear Tech node
	api, err := gsrpc.NewSubstrateAPI(url)
	require.NoError(t, err)

	b, err := hex.DecodeString("0c0842b0b24355170200")
	require.NoError(t, err)
	d := scale.NewDecoder(bytes.NewBuffer(b))

	blockHash, err := substrateTypes.NewHashFromHexString("0x6dffb564435e9878dca3a7faadbc72c6c6adbbfada02c75e8203623371e3d6bb")
	require.NoError(t, err)

	metadata, err := api.RPC.State.GetMetadata(blockHash)
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}

	factory := registry.NewFactory()
	eventRegistry, err := factory.CreateEventRegistry(metadata)
	require.NoError(t, err)

	var eventID substrateTypes.EventID

	err = d.Decode(&eventID)
	require.NoError(t, err)

	eventDecoder, ok := eventRegistry[eventID]
	require.True(t, ok)

	eventFields, err := eventDecoder.Decode(d)
	require.NoError(t, err)

	fmt.Println("Event ID: ", eventID)
	for _, field := range eventFields {
		fmt.Println(field)
	}

	// disp := &registry.Field{}
	// err = d.Decode(disp)
	// require.NoError(t, err)

	// byt, err := json.MarshalIndent(disp, "", "  ")
	// require.NoError(t, err)
	// fmt.Println(string(byt))

	require.True(t, false)
}
