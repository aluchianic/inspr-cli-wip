package hub

import (
	"github.com/spf13/cobra"
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"
)

// hubCommand represents the `hub` command
var hubCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	// rootCommand represents the base command when called without any subcommands
	hubCommand = &cobra.Command{
		Use:   "hub",
		Short: "Manage Inspr Hub",
	}
})

var _ = command.RegisterCommandInit(func() {
	cmd.RootCommandAddCommand(hubCommand)
})

func HubCommandAddCommand(c *cobra.Command) {
	hubCommand.AddCommand(c)
}
