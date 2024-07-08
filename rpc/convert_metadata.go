package rpc

import (
	"encoding/json"
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

type MetadataConverter struct {
	gearClients *firecoreRPC.Clients[*Client]
	logger      *zap.Logger
	tracer      logging.Tracer
}

func NewMetadataConverter(gearClients *firecoreRPC.Clients[*Client], logger *zap.Logger, tracer logging.Tracer) *MetadataConverter {
	return &MetadataConverter{
		gearClients: gearClients,
		logger:      logger,
		tracer:      tracer,
	}
}

func (mc *MetadataConverter) Convert(blockHash string) (string, error) {
	metadata, err := mc.fetchStateMetadata(blockHash)
	if err != nil {
		return "", fmt.Errorf("fetching state metadata: %w", err)
	}

	_ = metadata
	b, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshalling metadata: %w", err)
	}
	fmt.Println(string(b))

	return "", nil
}

func (mc *MetadataConverter) fetchStateMetadata(blockHash string) (*types.Metadata, error) {
	return firecoreRPC.WithClients(mc.gearClients, func(client *Client) (*types.Metadata, error) {
		if blockHash == "" {
			return client.state.GetMetadataLatest()
		}

		bHash, err := types.NewHashFromHexString(blockHash)
		if err != nil {
			return nil, fmt.Errorf("parsing block hash: %w", err)
		}

		fmt.Println("Fetching metadata at block hash", blockHash)
		return client.state.GetMetadata(bHash)
	})
}
