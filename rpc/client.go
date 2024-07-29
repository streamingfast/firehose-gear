package rpc

import (
	"fmt"

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
	eventsProvider      state.EventProvider
	eventsRetriever     retriever.EventRetriever
	state               rpcState.State
}

func NewClient(rpcEndpoint string) (*Client, error) {
	client, err := gsrpc.NewSubstrateAPI(rpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating API instance: %w", err)
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
		return nil, fmt.Errorf("error creating extrinsic retriever: %w", err)
	}

	eventProvider := state.NewEventProvider(client.RPC.State)
	eventRetriever, err := retriever.NewDefaultEventRetriever(eventProvider, client.RPC.State)
	if err != nil {
		return nil, fmt.Errorf("error creating event retriever: %w", err)
	}

	return &Client{
		client:              client,
		extrinsicsRetriever: extrinsicRetriever,
		eventsProvider:      eventProvider,
		eventsRetriever:     eventRetriever,
		state:               client.RPC.State,
	}, nil
}
