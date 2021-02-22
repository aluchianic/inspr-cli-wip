package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"
)

// authCommand represents the `auth` command
var authCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	authCommand = &cobra.Command{
		Use:     "auth",
		Aliases: []string{"register"},
		Short:   "Manage Inspr authorization",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Please see inspr auth --help for more commands\n")
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	cmd.RootCommandAddCommand(authCommand)
})

func AuthCommandAddCommand(cmd *cobra.Command) {
	authCommand.AddCommand(cmd)
}
