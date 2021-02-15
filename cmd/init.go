package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

var (
	apps          []string
	workspacePath string
	initCmd       = &cobra.Command{
		Use:   "init [workspaceName]",
		Args:  cobra.ExactArgs(1),
		Short: "[Workspace] Initialize Inspr workspace or dApp",
		Run: func(_ *cobra.Command, args []string) {
			workspace := configs.WorkspaceFiles{
				//Root: to change root path
			}
			err := workspace.Load()

			// Create workspace if not found
			if err != nil && err.NotFound() {
				err := workspace.Create(args[0], "workspace")
				configs.ShowAndExistIfErrorExists(err)
			}
			// Parse workspace to get config definition
			err = workspace.Parse()
			configs.ShowAndExistIfErrorExists(err)

			for _, app := range apps {
				err := workspace.Create(app, "application")
				configs.ShowAndExistIfErrorExists(err)
			}

		},
	}
)

// todo: init [workspace?] -a "apps-name"
func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringSliceVarP(&apps, "apps", "a", []string{}, "Init new Applications (should have also -w where to create)")
	initCmd.Flags().StringVarP(&workspacePath, "path", "p", "", "Path to workspace to be used, by default searching in current working dirrectory")
}
