package generator

import (
	"testing"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/rpc"
	"github.com/stretchr/testify/require"
)

func Test_Generator(t *testing.T) {
	gearClients := firecoreRPC.NewClients[*rpc.Client]()
	gearClient := rpc.NewClient("https://vara-mainnet.public.blastapi.io")
	gearClients.Add(gearClient)

	mc := rpc.NewMetadataConverter(gearClients, nil, nil)
	_, err := mc.Convert("")
	require.NoError(t, err)

	messages := mc.FetchMessages()
	metadata := mc.FetchMetadata()

	gen := NewGenerator("/Users/cbillett/devel/sf/firehose-gear/templates/gen_types.go.gotmpl", messages, metadata)
	err = gen.Generate()
	require.NoError(t, err)
}
