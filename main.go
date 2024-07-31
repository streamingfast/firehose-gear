package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	exec "github.com/centrifuge/go-substrate-rpc-client/v4/registry/exec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	retriever "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gobeam/stringy"
	gen_types "github.com/streamingfast/firehose-gear/templates"
)

func main() {
	//blockHash, err := types.NewHashFromHexString("0xdacec8ff307e047ff4dfa08a1a59dea0326b12902658e52de310a4214c04686c")
	blockHash, err := types.NewHashFromHexString("0xce88000c2153e31e367b800723aa3d78c537834e20439b2eb4c5db3870cd8e4c")
	//blockHash, err := types.NewHashFromHexString("0xafd423f4bebb20864e535197b732c343ba72bb446574c9ceb117bf4f721098da")
	if err != nil {
		log.Fatalf("Failed to create block hash: %v", err)
	}

	url := "https://vara-mainnet.public.blastapi.io" // Replace with the actual URL of your Gear Tech node
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		log.Fatal("Error creating API instance")
	}

	block, err := api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		log.Fatalf("Failed to get block: %v", err)
	}

	j, err := json.MarshalIndent(block, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal block: %v", err)
	}
	_ = j
	// fmt.Println("block: ", string(j))

	chain := generic.NewDefaultChain(api.Client)
	factory := registry.NewFactory()

	extrinsicParser := parser.NewDefaultExtrinsicParser()
	extrinsicRetriever, err := retriever.NewDefaultExtrinsicRetriever(
		extrinsicParser,
		chain,
		api.RPC.State,
		factory,
		exec.NewRetryableExecutor[*generic.DefaultGenericSignedBlock](),
		exec.NewRetryableExecutor[[]*parser.DefaultExtrinsic](),
	)

	if err != nil {
		log.Fatalf("Error creating extrinsic retriever: %v", err)
	}

	extrinsics, err := extrinsicRetriever.GetExtrinsics(blockHash)
	if err != nil {
		log.Fatalf("Failed to get extrinsics: %v", err)
	}

	metadata, err := api.RPC.State.GetMetadata(blockHash)
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}
	_ = metadata

	for _, extrinsic := range extrinsics {
		parts := strings.Split(extrinsic.Name, ".")
		pallet := parts[0]
		call := parts[1]
		call = stringy.New(call).PascalCase().Get()

		structName := pallet + "_" + call + "Call"
		funcName := "To_" + structName

		if fn, found := gen_types.FuncMap[funcName]; found {
			o := fn.Call([]reflect.Value{reflect.ValueOf(extrinsic.CallFields)})
			fmt.Println(funcName+":", o)
			continue
		}

		fmt.Println("WTF:", funcName)

	}
}
