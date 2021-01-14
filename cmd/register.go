package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register account to get token for Inspr",
	Long:  "Prompt: (acc, pass) -> server -> token",
	Run:   func(cmd *cobra.Command, args []string) {},
}
