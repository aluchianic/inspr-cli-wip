package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/configs"
	"os"
)

var (
	app     string
	path    string
	initCmd = &cobra.Command{
		Use:   "init [workspaceName]",
		Args:  cobra.ExactArgs(1),
		Short: "[Workspace] Initialize Inspr workspace or dApp",
		Run: func(_ *cobra.Command, args []string) {
			wName := args[0]

			_, err := configs.InitWorkspace(path)
			if err != nil && err.NotFound() {
				if err := configs.CreateWorkspace(wName); err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
					os.Exit(1)
				}
			}
			err = configs.DescribeWorkspace()

			if app != "" {
				_, err := configs.InitApp(app)
				if err != nil && err.NotFound() {
					if err := configs.CreateApp(app); err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
						os.Exit(1)
					}
				}
				err = configs.DescribeApp()
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
