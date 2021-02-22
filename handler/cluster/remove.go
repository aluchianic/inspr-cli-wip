package cluster

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

var (
	// removeCmd represents the `cluster remove` command
	removeCmd *cobra.Command
	force     bool
)

var _ = command.RegisterCommandVar(func() {
	removeCmd = &cobra.Command{
		Use:   "remove [application name]",
		Short: "Remove deployed dApp from cluster",
		Args:  cobra.MinimumNArgs(1),
		Long:  "By default command tries to remove in 'safe' mode, awaiting to close all connections",
		Run:   func(cmd *cobra.Command, appNames []string) {},
	}
})

var _ = command.RegisterCommandInit(func() {
	ClusterAddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Remove dApp without graceful shutdown of service.")
	// --safe [--default]
})
