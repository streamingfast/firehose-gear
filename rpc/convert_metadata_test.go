package rpc

import (
	"testing"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/stretchr/testify/require"
)

func Test_ConvertMetadata(t *testing.T) {
	gearClients := firecoreRPC.NewClients[*Client]()
	gearClient, err := NewClient("https://vara-mainnet.public.blastapi.io")
	require.NoError(t, err)
	gearClients.Add(gearClient)

	mc := NewMetadataConverter(gearClients, nil, nil)
	err = mc.Convert("")
	if err != nil {
		println("")
	}
	require.NoError(t, err)
}
