package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

// Private helpers ------------------------------------------------------------

func command() *cobra.Command {
	return &cobra.Command{
		Use:     "cli",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Hello from cli")
		},
	}
}

func main() {
	if err := command().Execute(); err != nil {
		os.Exit(1)
	}
}
