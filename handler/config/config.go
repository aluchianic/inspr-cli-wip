package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"
)

var (
	// configCommand represents the `config` command
	configCommand *cobra.Command
	appNames      []string
	workspacePath string
)

var _ = command.RegisterCommandVar(func() {
	configCommand = &cobra.Command{
		Use:     "config",
		Aliases: []string{"init"},
		Short:   "Manage Inspr configs",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Printf("Please see inspr config --help for more commands\n")
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	cmd.RootCommandAddCommand(configCommand)

	appCommand.Flags().StringSliceVarP(&appNames, "appNames", "a", []string{}, "Init new Applications (should have also -w where to create)")
	appCommand.Flags().StringVarP(&workspacePath, "path", "p", "", "Path to workspace to be used, by default searching in current working directory")
})

func ConfigCommandAddCommand(c *cobra.Command) {
	configCommand.AddCommand(c)
}
