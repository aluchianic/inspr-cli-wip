package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	all bool
)

func init() {
	rootCmd.AddCommand(statsCmd)

	statsCmd.Flags().BoolVarP(&all, "all", "a", false, "Return statistics from all clusters")
}

var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"list"},
	Short:   "[Cluster] Return statistics from clusters, by default returns statistics only for running clusters",
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			fmt.Println("Return statistics for ALL clusters ...")
		} else {
			fmt.Println("Return statistics for current running clusters ...")
		}

	},
}
