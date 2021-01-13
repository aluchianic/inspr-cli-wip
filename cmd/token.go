package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(tokenCmd)

	statsCmd.Flags().StringP("create", "c", "", "Create new token.")
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Create or validate Inpsr token",
	Run:   func(cmd *cobra.Command, args []string) {},
}
