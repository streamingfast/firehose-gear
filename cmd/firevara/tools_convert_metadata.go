package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	firecore "github.com/streamingfast/firehose-core"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/rpc"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewToolsConvertMetadataCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert <endpoint>",
		Short: "Fetch latest or specified metadata from blockhash",
		RunE:  convertMetadataE(logger, tracer),
	}

	cmd.Flags().String("block-hash", "", "Blockhash, prefixed by 0x, to fetch metadata from, leave empty for latest")

	return cmd
}

func convertMetadataE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) error {
		blockHash := sflags.MustGetString(cmd, "block-hash")
		endpoint := args[0]

		if endpoint == "" {
			return fmt.Errorf("missing endpoint")
		}

		if blockHash != "" && !strings.HasPrefix(blockHash, "0x") {
			blockHash = "0x" + blockHash
		}

		gearClients := firecoreRPC.NewClients[*rpc.Client]()
		gearClient := rpc.NewClient(endpoint)
		gearClients.Add(gearClient)

		metadataConverter := rpc.NewMetadataConverter(gearClients, logger, tracer)
		metadata, err := metadataConverter.Convert(blockHash)
		if err != nil {
			return fmt.Errorf("converting metadata: %w", err)
		}

		for _, t := range metadata.Types {
			msg := t.ToProtoMessage()
			if msg == "" {
				continue
			}
			fmt.Println(msg)
		}

		// fmt.Println(metadata)
		return nil
	}
}
