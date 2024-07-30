package main

import "github.com/spf13/cobra"

var GenerateCmd = &cobra.Command{Use: "generate", Short: "Generate decoded block or types from metadata"}

func init() {
	GenerateCmd.AddCommand(NewToolsGenerateDecodedBlock(logger, tracer))
}
