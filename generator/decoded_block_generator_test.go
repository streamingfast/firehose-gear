package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DecodedBlockGenerator(t *testing.T) {
	// TODO: here we want to take in a block saved locally, decode all the fields
	// and then generate the new block with the decoded fields
	dbg := NewDecodedBlockGenerator("../templates/decoded_block.proto.gotmpl")
	err := dbg.Generate()
	require.NoError(t, err)
}
