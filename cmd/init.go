package cmd

import (
	"github.com/spf13/cobra"
	c "inspr-cli/configs"
)

var (
	app     string
	initCmd = &cobra.Command{
		Use:   "init [workspace?]",
		Args:  cobra.MaximumNArgs(1),
		Short: "[Workspace] Initialize Inspr workspace or dApp",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) > 0 {
				c.SetWorkspaceName(args[0])
			}

			if found := c.LoadWorkspaceConfig(); !found {
				c.WriteWorkspaceConfig()
			}

			if app != "" {
				c.SetAppName(app)
				if found := c.LoadAppConfig(); !found {
					c.WriteAppConfigToDisk()
				}
			}
		},
	}
)

// todo: init [workspace?] -a "app-name"
func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&app, "app", "a", "", "Init new DApp (should have also -w where to create)")
}
