package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
)

var (
	// configCommand represents the `config` command
	configCommand *cobra.Command
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
	config.CM().Flags.AddFlags(configCommand.PersistentFlags())
})

func ConfigCommandAddCommand(c *cobra.Command) {
	configCommand.AddCommand(c)
}
