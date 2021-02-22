package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

var (
	// statsCmd represents the `cluster stats` command
	statsCmd *cobra.Command
	full     bool
)

var _ = command.RegisterCommandVar(func() {
	statsCmd = &cobra.Command{
		Use:     "stats",
		Aliases: []string{"list"},
		Short:   "Return statistics from clusters, by default returns statistics only for running clusters",
		Run: func(cmd *cobra.Command, args []string) {
			if full {
				fmt.Println("Return statistics for ALL clusters ...")
			} else {
				fmt.Println("Return statistics for current running clusters ...")
			}

		},
	}
})

var _ = command.RegisterCommandInit(func() {
	ClusterAddCommand(statsCmd)
	statsCmd.Flags().BoolVarP(&full, "full", "f", false, "Return statistics from all clusters")
})
