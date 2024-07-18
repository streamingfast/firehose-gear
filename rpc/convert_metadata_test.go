package rpc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

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
	b, err := hex.DecodeString("0c0842b0b24355170200")
	require.NoError(t, err)
	d := scale.NewDecoder(bytes.NewBuffer(b))

	disp := &substrateTypes.DispatchInfo{}

	err = d.Decode(disp)
	require.NoError(t, err)

	byt, err := json.MarshalIndent(disp, "", "  ")
	require.NoError(t, err)
	fmt.Println(string(byt))

	require.True(t, false)
}
