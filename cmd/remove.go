package cmd

import (
	"github.com/spf13/cobra"
	c "inspr-cli/configs"
)

var force bool

func init() {
	rootCmd.AddCommand(removeCmd)

	c.AddWorkspaceFlag(removeCmd)
	c.AddTagFlag(removeCmd)

	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Remove dApp without graceful shutdown of service.")
	// --safe [--default]
}

var removeCmd = &cobra.Command{
	Use:   "remove [id...]",
	Short: "[Cluster] Remove deployed dApp from cluster",
	Long:  "By default command tries to remove in 'safe' mode, awaiting to close all connections",
	Run: func(cmd *cobra.Command, ids []string) {
		found := c.LoadWorkspaceConfig()
		if !found {
			panic("No workspace found.")
		}
		c.LoadAppConfig()
	},
}
