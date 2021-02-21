package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"
)

// clusterCmd represents the cluster command
var clusterCmd *cobra.Command

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
})

func ClusterAddCommand(cmd *cobra.Command) {
	clusterCmd.AddCommand(cmd)
}
