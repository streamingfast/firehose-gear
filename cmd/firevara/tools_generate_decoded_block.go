package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	firecore "github.com/streamingfast/firehose-core"
	firecoreRPC "github.com/streamingfast/firehose-core/rpc"
	"github.com/streamingfast/firehose-gear/generator"
	"github.com/streamingfast/firehose-gear/rpc"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewToolsGenerateExtrinsicProto(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extrinsic-proto <endpoint>",
		Short: "Fetch the latest metadata from a block and generate the extrinsic proto",
		RunE:  generateDecodedBlockE(logger, tracer),
	}

	cmd.Flags().String("blockhash", "", "Blockhash, prefixed by 0x, to fetch metadata from, leave empty for latest")
	cmd.Flags().String("output", "proto/sf/gear/extrinsic/type/v1/extrinsic.proto", "Output extrinsic file location")

	return cmd
}

func generateDecodedBlockE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) error {
		blockhash := sflags.MustGetString(cmd, "blockhash")
		outputPath := sflags.MustGetString(cmd, "output")
		endpoint := args[0]

		if endpoint == "" {
			return fmt.Errorf("missing endpoint")
		}

		if blockhash != "" && !strings.HasPrefix(blockhash, "0x") {
			blockhash = "0x" + blockhash
		}

		logger.Info("generating decoded block", zap.String("endpoint", endpoint), zap.String("blockhash", blockhash), zap.String("output_path", outputPath))

		gearClients := firecoreRPC.NewClients[*rpc.Client]()
		gearClient, err := rpc.NewClient(endpoint)
		if err != nil {
			return fmt.Errorf("error creating gear client: %w", err)
		}
		gearClients.Add(gearClient)

		metadataConverter := rpc.NewMetadataConverter(gearClients, logger, tracer)
		err = metadataConverter.Convert(blockhash)
		if err != nil {
			return fmt.Errorf("converting metadata: %w", err)
		}

		messages := metadataConverter.FetchMessages()

		// sort the messages to have them in the same order on each run
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].FullTypeName() < messages[j].FullTypeName()
		})

		dbg := generator.NewDecodedBlockGenerator("templates/extrinsic.proto.gotmpl", messages)
		content, err := dbg.Generate()
		if err != nil {
			return fmt.Errorf("generating decoded block: %w", err)
		}

		err = os.MkdirAll("proto/sf/gear/extrinsic/type/v1", os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		err = os.WriteFile(outputPath, content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		return nil
	}
}
