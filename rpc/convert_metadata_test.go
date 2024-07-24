package rpc

import (
	"testing"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/stretchr/testify/require"
)

func Test_ConvertMetadata(t *testing.T) {
	gearClients := firecoreRPC.NewClients[*Client]()
	gearClient := NewClient("https://vara-mainnet.public.blastapi.io")
	gearClients.Add(gearClient)

	mc := NewMetadataConverter(gearClients, nil, nil)
	err := mc.Convert("")
	require.NoError(t, err)
}
