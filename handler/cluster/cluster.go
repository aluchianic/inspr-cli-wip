package cluster

import (
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"

	"fmt"
	"github.com/spf13/cobra"
)

// clusterCmd represents the cluster command
var (
	clusterCmd *cobra.Command
)

var _ = command.RegisterCommandVar(func() {
	clusterCmd = &cobra.Command{
		Use:     "cluster",
		Aliases: []string{"clusters"},
		Short:   "Manage Inspr cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Please see inspr cluster --help for more commands\n")
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	cmd.RootCommandAddCommand(clusterCmd)
	config.CM().Flags.AddFlags(clusterCmd.PersistentFlags())
})

func ClusterAddCommand(cmd *cobra.Command) {
	clusterCmd.AddCommand(cmd)
}
