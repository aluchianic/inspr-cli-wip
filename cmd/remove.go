package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)

	deployCommand.Flags().BoolP("force", "f", false, "Remove dApp even if other resources depends on it.")
}

var removeCmd = &cobra.Command{
	Use:   "remove [id...]",
	Short: "[Cluster] Remove deployed dApp from cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}
