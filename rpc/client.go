package rpc

import (
	"log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	exec "github.com/centrifuge/go-substrate-rpc-client/v4/registry/exec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	retriever "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/chain/generic"
	rpcState "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
)

type Client struct {
	client              *gsrpc.SubstrateAPI
	extrinsicsRetriever retriever.DefaultExtrinsicRetriever
	eventsRetriever     retriever.EventRetriever
	state               rpcState.State
}

func NewClient(rpcEndpoint string) *Client {
	client, err := gsrpc.NewSubstrateAPI(rpcEndpoint)
	if err != nil {
		log.Fatal("Error creating API instance")
	}
	chain := generic.NewDefaultChain(client.Client)
	factory := registry.NewFactory()

	extrinsicParser := parser.NewDefaultExtrinsicParser()
	extrinsicRetriever, err := retriever.NewDefaultExtrinsicRetriever(
		extrinsicParser,
		chain,
		client.RPC.State,
		factory,
		exec.NewRetryableExecutor[*generic.DefaultGenericSignedBlock](),
		exec.NewRetryableExecutor[[]*parser.DefaultExtrinsic](),
	)
	if err != nil {
		log.Fatalf("Error creating extrinsic retriever: %v", err)
	}

	eventRetriever, err := retriever.NewDefaultEventRetriever(state.NewEventProvider(client.RPC.State), client.RPC.State)
	if err != nil {
		log.Fatalf("Error creating event retriever: %v", err)
	}

	return &Client{
		client:              client,
		extrinsicsRetriever: extrinsicRetriever,
		eventsRetriever:     eventRetriever,
		state:               client.RPC.State,
	}
}
