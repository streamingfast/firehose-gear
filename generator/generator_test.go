package generator

import (
	"sort"
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

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].FullTypeName() < messages[j].FullTypeName()
	})

	gen := NewGenerator("../templates/gen_types.go.gotmpl", messages, metadata, mc.TypeConverter.IdToField)
	err = gen.Generate()
	require.NoError(t, err)
}
