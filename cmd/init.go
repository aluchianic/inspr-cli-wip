package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

var (
	app     string
	initCmd = &cobra.Command{
		Use:   "init [workspace?]",
		Args:  cobra.MaximumNArgs(1),
		Short: "[Workspace] Initialize Inspr workspace or dApp",
		Run: func(_ *cobra.Command, args []string) {
			w := configs.Workspace{}
			if len(args) > 0 {
				w.Name = args[0]
			}
			if !w.Init() {
				w.Create()
			}

			if app != "" {
				a := configs.App{Name: app}
				if !a.Init() {
					a.Create()
				}
			}
		},
	}
)

// todo: init [workspace?] -a "app-name"
func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&app, "app", "a", "", "Init new AppConfig (should have also -w where to create)")
}
