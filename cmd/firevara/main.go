package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "firevara",
	Short: "Firehose Gear (for Vara) block fetching and tooling",
	Args:  cobra.ExactArgs(1),
}

func init() {
	logging.InstantiateLoggers(logging.WithDefaultLevel(zap.InfoLevel))

	rootCmd.AddCommand(NewFetchCmd(logger, tracer))
	rootCmd.AddCommand(ToolsCmd)
	rootCmd.AddCommand(GenerateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
