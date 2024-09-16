package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/spf13/cobra"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewToolsFetchMetadataCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tools-fetch-metadata",
		Short: "Fetch all the metadata for Vara blockchain with all the specVersion defined in the file tools-data/specVersions.txt",
		RunE:  toolsFetchMetadataRunE(logger, tracer),
	}

	cmd.Flags().String("spec-version-path", "", "Spec Versions text file, format: <version> <blockhash>")

	return cmd
}

type SpecVerion struct {
	SpecVersion string
	BlockHash   string
}

func toolsFetchMetadataRunE(logger *zap.Logger, _ logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) (err error) {
		filePath := cmd.Flag("spec-version-path").Value.String()
		readFile, err := os.Open(filePath)

		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		defer readFile.Close()

		specVersions := make([]*SpecVerion, 0)
		for _, line := range fileLines {
			l := strings.Split(line, " ")
			specVersions = append(specVersions, &SpecVerion{
				SpecVersion: l[0],
				BlockHash:   l[1],
			})
		}

		url := "https://vara-mainnet.public.blastapi.io" // Replace with the actual URL of your Gear Tech node
		api, err := gsrpc.NewSubstrateAPI(url)

		for _, sv := range specVersions {
			logger.Debug("fetching metadata for spec version", zap.String("version", sv.SpecVersion), zap.String("block_hash", sv.BlockHash))
			blockhash, err := types.NewHashFromHexString(sv.BlockHash)
			if err != nil {
				return fmt.Errorf("failed to convert block hash: %w", err)
			}

			metadata, err := api.RPC.State.GetMetadata(blockhash)
			if err != nil {
				return fmt.Errorf("failed to get metadata: %w", err)
			}

			b, err := json.Marshal(metadata)
			if err != nil {
				return fmt.Errorf("failed to marshal metadata: %w", err)
			}

			err = os.WriteFile(fmt.Sprintf("%s.json", sv.SpecVersion), b, 0644)
			if err != nil {
				return fmt.Errorf("failed to write metadata: %w", err)
			}
		}

		return
	}
}
