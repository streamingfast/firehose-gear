package rpc

import (
	"fmt"
	"testing"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/stretchr/testify/require"
)

func Test_ConvertMetadata(t *testing.T) {
	gearClients := firecoreRPC.NewClients[*Client]()
	gearClient := NewClient("https://vara-mainnet.public.blastapi.io")
	gearClients.Add(gearClient)

	mc := NewMetadataConverter(gearClients, nil, nil)
	types, err := mc.Convert("")
	require.NoError(t, err)

	for _, tt := range types.Types {
		fmt.Println(tt.ToProtoMessage())
	}
	require.True(t, false)
}
