package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

var (
	force bool
	w     string
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Remove dApp without graceful shutdown of service.")
	removeCmd.Flags().StringVarP(&w, "workspace", "w", "", "Set path to workspace")
	// --safe [--default]
}

var removeCmd = &cobra.Command{
	Use:   "remove [id...]",
	Short: "[Cluster] Remove deployed dApp from cluster",
	Long:  "By default command tries to remove in 'safe' mode, awaiting to close all connections",
	Run: func(cmd *cobra.Command, ids []string) {
		_ = configs.Workspace{Path: w}
		a := configs.App{
			Id: ids[0],
		}
		a.Describe()
	},
}
