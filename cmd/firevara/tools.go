package main

import (
	"github.com/spf13/cobra"
)

var ToolsCmd = &cobra.Command{Use: "tools", Short: "Developer tools for operators and developers"}

func init() {
	ToolsCmd.AddCommand(NewToolsConvertMetadataCmd(logger, tracer))
}
