package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/configs"
	"path"
	"strings"
)

var (
	apps    []string
	_path   string
	initCmd = &cobra.Command{
		Use:   "init [workspaceName]",
		Args:  cobra.ExactArgs(1),
		Short: "[Workspace] Initialize Inspr workspace or dApp",
		Run: func(_ *cobra.Command, args []string) {
			workspace := configs.WorkspaceFiles{}

			err := workspace.Load("workspace")
			// Create workspace if not found
			if err != nil && err.NotFound() {
				err := workspace.Create(args[0])
				configs.ShowAndExistIfErrorExists(err)
			}
			// Parse workspace to get config definition
			workspace.Parse()

			for _, app := range apps {
				p := strings.ReplaceAll(workspace.Path, path.Base(workspace.Path), "")

				f := configs.FileRaw{
					Path:       path.Join(p, "apps"),
					Definition: "application",
				}
				err := f.Create(app)
				configs.ShowAndExistIfErrorExists(err)
			}
		},
	}
)

// todo: init [workspace?] -a "apps-name"
func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringSliceVarP(&apps, "apps", "a", []string{}, "Init new Applications (should have also -w where to create)")
	initCmd.Flags().StringVarP(&_path, "path", "p", "", "Path to workspace to be used, by default searching in current working dirrectory")
}
