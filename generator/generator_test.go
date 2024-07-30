package generator

import (
	"fmt"
	"sort"
	"testing"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/rpc"
	"github.com/stretchr/testify/require"
)

func Test_Generator(t *testing.T) {

	fmt.Println("Test_Generator")
	gearClients := firecoreRPC.NewClients[*rpc.Client]()
	gearClient, err := rpc.NewClient("https://vara-mainnet.public.blastapi.io")
	require.NoError(t, err)
	gearClients.Add(gearClient)

	mc := rpc.NewMetadataConverter(gearClients, nil, nil)
	err = mc.Convert("")
	require.NoError(t, err)

	messages := mc.FetchMessages()
	metadata := mc.FetchMetadata()

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].FullTypeName() < messages[j].FullTypeName()
	})

	gen := NewGenerator("../templates/gen_types.go.gotmpl", messages, metadata)
	err = gen.Generate()
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(t, err)
	fmt.Println("All done!")
}
