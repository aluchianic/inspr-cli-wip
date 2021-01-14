package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hubCmd)
}

var hubCmd = &cobra.Command{
	Use:   "hub",
	Short: "Commands to operate with Inpsr HUB",
}
