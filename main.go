package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	gen_types "github.com/streamingfast/firehose-gear/templates"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	exec "github.com/centrifuge/go-substrate-rpc-client/v4/registry/exec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	retriever "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gobeam/stringy"
)

func main() {
	blockHash, err := types.NewHashFromHexString("0xad1ce65fb2dde8ee826f2b25999f873f92815224300cd3b747031123f4cb6611")
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

	// metadata, err := api.RPC.State.GetMetadata(blockHash)
	//if err != nil {
	//	log.Fatalf("Failed to get metadata: %v", err)
	//}
	//
	for _, extrinsic := range extrinsics {
		parts := strings.Split(extrinsic.Name, ".")
		pallet := parts[0]
		call := parts[1]
		call = stringy.New(call).PascalCase().Get()
		structName := pallet + "_" + call + "_Call"

		switch structName {
		case "Timestamp_Set_Call":
			fmt.Println(call)
		case "Balances_TransferKeepAlive_Call":
			gen_types.To_Balances_TransferKeepAliveCall(extrinsic.CallFields)
			fmt.Println(call)
		}
	}
}
