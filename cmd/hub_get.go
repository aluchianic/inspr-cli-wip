package cmd

import "github.com/spf13/cobra"

func init() {
	hubCmd.AddCommand(hubGetCmd)
}

var hubGetCmd = &cobra.Command{
	Use:   "get [dAppName...]",
	Short: "Get dApp from Inspr Hub",
	Run:   func(cmd *cobra.Command, ids []string) {},
}
