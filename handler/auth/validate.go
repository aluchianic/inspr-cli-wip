package auth

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

// validateCommand represents the `auth validate` command
var validateCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	validateCommand = &cobra.Command{
		Use:   "validate",
		Short: "Validate your Inspr authorization",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
})

var _ = command.RegisterCommandInit(func() {
	AuthCommandAddCommand(validateCommand)
})
