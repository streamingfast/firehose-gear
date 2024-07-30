package generator

import (
	"os"
	"sort"
	"testing"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/rpc"
	"github.com/stretchr/testify/require"
)

func Test_DecodedBlockGenerator(t *testing.T) {
	gearClients := firecoreRPC.NewClients[*rpc.Client]()
	gearClient, err := rpc.NewClient("https://vara-mainnet.public.blastapi.io")
	require.NoError(t, err)
	gearClients.Add(gearClient)

	mc := rpc.NewMetadataConverter(gearClients, nil, nil)
	err = mc.Convert("")
	require.NoError(t, err)

	messages := mc.FetchMessages()

	// sort the messages to have them in the same order on each run
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].FullTypeName() < messages[j].FullTypeName()
	})

	dbg := NewDecodedBlockGenerator("../templates/decoded_block.proto.gotmpl", messages)
	content, err := dbg.Generate()
	require.NoError(t, err)

	err = os.WriteFile("../templates/decoded_block.proto", content, 0644)
	require.NoError(t, err)
}
