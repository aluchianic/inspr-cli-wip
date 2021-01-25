package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

var force bool

func init() {
	rootCmd.AddCommand(removeCmd)

	configs.AddPathFlag(removeCmd)

	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Remove dApp without graceful shutdown of service.")
	// --safe [--default]
}

var removeCmd = &cobra.Command{
	Use:   "remove [appName...]",
	Short: "[Cluster] Remove deployed dApp from cluster",
	Args:  cobra.MinimumNArgs(1),
	Long:  "By default command tries to remove in 'safe' mode, awaiting to close all connections",
	Run:   func(cmd *cobra.Command, appNames []string) {},
}
