package cluster

import (
	"inspr-cli/pkg/command"

	"github.com/spf13/cobra"
)

// describeCommand represents the `cluster describe` command
var describeCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	describeCommand = &cobra.Command{
		Use:   "describe [application name]",
		Short: "DescribeApp dApp with it dependencies, channel types and third parties",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	}
})

var _ = command.RegisterCommandInit(func() {
	ClusterAddCommand(describeCommand)
})
