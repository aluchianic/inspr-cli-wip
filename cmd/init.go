package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

var (
	app     string
	path    string
	initCmd = &cobra.Command{
		Use:   "init [workspaceName]",
		Args:  cobra.ExactArgs(1),
		Short: "[Workspace] Initialize Inspr workspace or dApp",
		Run: func(_ *cobra.Command, args []string) {
			var (
				wName = args[0]
				w     = configs.NewWorkspace(wName)
				a     *configs.Application
				err   *configs.ConfigError
			)

			if app != "" {
				a, err = w.NewApplication(app)
				if err != nil && err.AlreadyExists() {
					fmt.Printf("Application `%s` already exists \n", app)
				} else {
					configs.ShowAndExistIfErrorExists(err)
				}
				err = a.Describe()
				configs.ShowAndExistIfErrorExists(err)
			}
		},
	}
)

// todo: init [workspace?] -a "app-name"
func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&app, "app", "a", "", "Init new AppConfig (should have also -w where to create)")
	initCmd.Flags().StringVarP(&path, "path", "p", "", "Path to workspace to be used, by default searching in current working dirrectory")
}
