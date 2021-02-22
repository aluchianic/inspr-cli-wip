package hub

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

// hubGetCmd represents the `hub get` command
var hubGetCmd *cobra.Command

var _ = command.RegisterCommandVar(func() {
	// rootCommand represents the base command when called without any subcommands
	hubGetCmd = &cobra.Command{
		Use:   "get [application names...]",
		Short: "Get applications from Inspr Hub",
		Run:   func(cmd *cobra.Command, ids []string) {},
	}
})

var _ = command.RegisterCommandInit(func() {
	HubCommandAddCommand(hubGetCmd)
})
