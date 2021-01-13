package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(graphCmd)
}

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "[dApp] Display dApp information",
	Run:   func(cmd *cobra.Command, args []string) {},
}
