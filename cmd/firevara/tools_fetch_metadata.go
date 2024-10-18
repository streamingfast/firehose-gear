package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
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
	cmd.Flags().StringArray("versions", []string{"all"}, "Spec Versions to save locally, format: <version>.json")
	cmd.Flags().String("output-dir", "metadata", "Output directory to save the metadata")

	return cmd
}

type SpecVerion struct {
	SpecVersion string
	BlockHash   string
}

func toolsFetchMetadataRunE(logger *zap.Logger, _ logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) (err error) {
		filePath := sflags.MustGetString(cmd, "spec-version-path")
		logger.Debug("spec version file path", zap.String("path", filePath))
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
		versionsToSave := sflags.MustGetStringArray(cmd, "versions")
		logger.Debug("spec versions to save", zap.Strings("versions", versionsToSave))
		saveAll := false
		if len(versionsToSave) == 1 && versionsToSave[0] == "all" {
			saveAll = true
		}

		outputDir := sflags.MustGetString(cmd, "output-dir")
		logger.Debug("output directory", zap.String("dir", outputDir))
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			err = os.Mkdir(outputDir, 0755)
			if err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}
		}

		for _, sv := range specVersions {
			if saveAll || slices.Contains(versionsToSave, sv.SpecVersion) {
				logger.Debug("fetching metadata for spec version", zap.String("version", sv.SpecVersion), zap.String("block_hash", sv.BlockHash))
				blockhash, err := types.NewHashFromHexString(sv.BlockHash)
				if err != nil {
					return fmt.Errorf("failed to convert block hash: %w", err)
				}

				metadata, err := api.RPC.State.GetMetadata(blockhash)
				if err != nil {
					return fmt.Errorf("failed to get metadata: %w", err)
				}

				b, err := json.MarshalIndent(metadata, "", "  ")
				if err != nil {
					return fmt.Errorf("failed to marshal metadata: %w", err)
				}

				filename := fmt.Sprintf("metadata%s.json", sv.SpecVersion)
				dst, err := os.Create(filepath.Join(outputDir, filename))
				if err != nil {
					return fmt.Errorf("failed to create metadata file: %w", err)
				}

				_, err = dst.WriteString(string(b))
				if err != nil {
					return fmt.Errorf("failed to write metadata file: %w", err)
				}
			}
		}

		return
	}
}
