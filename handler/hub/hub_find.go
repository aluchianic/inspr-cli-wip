package hub

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

// hubFindCommand represents the `hub find` command
var hubFindCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	hubFindCommand = &cobra.Command{
		Use:   "find [application name]",
		Short: "Find application in Inspr Hub",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, dAppName []string) {},
	}
})

var _ = command.RegisterCommandInit(func() {
	HubCommandAddCommand(hubFindCommand)
})
