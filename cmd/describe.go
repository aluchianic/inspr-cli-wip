package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(describeCmd)
}

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe dApp with it dependencies, channel types and third parties",
	Run:   func(cmd *cobra.Command, args []string) {},
}
