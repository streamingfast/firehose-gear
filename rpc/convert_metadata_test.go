package rpc

import (
	"fmt"
	"testing"

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

	protobufMessages := &types.ProtobufMessages{}

	for key := range ttypes {
		protobufMessages.Messages = append(protobufMessages.Messages, &types.ProtobufMessage{
			Name: key,
		})
	}

	fmt.Println(protobufMessages.ToProtoMessage())
	require.True(t, false)
}
