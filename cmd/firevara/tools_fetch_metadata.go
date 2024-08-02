package main

import (
	"bufio"
	"fmt"
	"os"

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

	return cmd
}

func toolsFetchMetadataRunE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) (err error) {
		// curl --location 'https://vara-mainnet.public.blastapi.io' --header 'Content-Type: application/json' --data '{ "id": 1, "jsonrpc": "2.0", "method": "state_getMetadata", "params": ["0x68a1e81e50b4bd3c7bc35727eca22d31cd68808294d71863bbfac8da38514cf4"] }' | jq .result
		filePath := "/Users/eduardvoiculescu/git/streamingfast/firehose-gear/cmd/firebara/tools-data/specVersions.txt"
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

		readFile.Close()

		for _, line := range fileLines {
			fmt.Println(line)
		}

		return
	}
}
