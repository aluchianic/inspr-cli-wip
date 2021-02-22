package auth

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

// authCommand represents the `auth register` command
var registerCmd *cobra.Command

var _ = command.RegisterCommandVar(func() {
	// rootCommand represents the base command when called without any subcommands
	registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register account to get Inspr authorization",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
})

var _ = command.RegisterCommandInit(func() {
	AuthCommandAddCommand(registerCmd)
})
