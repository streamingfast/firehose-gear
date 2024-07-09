package main

import (
	"encoding/json"
	"fmt"
	"log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	exec "github.com/centrifuge/go-substrate-rpc-client/v4/registry/exec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	retriever "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func main() {
	blockHash, err := types.NewHashFromHexString("0x53cde3f3c1909cdc98c4e724c569f5517511928d451811916ebb181bc7d7f7f7")
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

	// fmt.Println("---------------------------------------------")
	// fmt.Println("- Extrinsics")
	// fmt.Println("---------------------------------------------")
	// fmt.Println("Number of extrinsics: ", len(extrinsics))
	for _, extrinsic := range extrinsics {
		j, err := json.MarshalIndent(extrinsic, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal extrinsic: %v", err)
		}
		_ = j
		// fmt.Println(string(j))
	}

	eventRetriever, err := retriever.NewDefaultEventRetriever(state.NewEventProvider(api.RPC.State), api.RPC.State)
	if err != nil {
		log.Fatalf("Error creating event retriever: %v", err)
	}

	events, err := eventRetriever.GetEvents(blockHash)
	if err != nil {
		log.Fatalf("Failed to get events: %v", err)
	}

	// fmt.Println("---------------------------------------------")
	// fmt.Println("- Events")
	// fmt.Println("---------------------------------------------")
	// fmt.Println("Number of events: ", len(events))

	for _, event := range events {
		j, err := json.MarshalIndent(event, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal event: %v", err)
		}
		_ = j

		// fmt.Println(string(j))
		//log.Printf("Event ID: %x \n", event.EventID)
		//log.Printf("Event Name: %s \n", event.Name)
		//log.Printf("Event Fields Count: %d \n", len(event.Fields))
		//for k, v := range event.Fields {
		//	log.Printf("Field Name: %s \n", k)
		//	log.Printf("Field Type: %v \n", reflect.TypeOf(v))
		//	log.Printf("Field Value: %v \n", v)
		//}
	}

	metadata, err := api.RPC.State.GetMetadata(blockHash)
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}

	m, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal extrinsic: %v", err)
	}

	_ = m

	block0, err := types.NewHashFromHexString("0x1555a60f789e322c607f8b8df062cf0441484897b547028c378509d5e56eec48")
	if err != nil {
		log.Fatalf("Failed to create block hash: %v", err)
	}
	specVerion, err := api.RPC.State.GetRuntimeVersion(block0)
	if err != nil {
		log.Fatalf("Failed to get runtime version: %v", err)
	}

	fmt.Println("Spec Version: ", specVerion.SpecVersion)
}
