package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	full bool
)

func init() {
	rootCmd.AddCommand(statsCmd)

	statsCmd.Flags().BoolVarP(&full, "full", "f", false, "Return statistics from all clusters")
}

var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"list"},
	Short:   "[Cluster] Return statistics from clusters, by default returns statistics only for running clusters",
	Run: func(cmd *cobra.Command, args []string) {
		if full {
			fmt.Println("Return statistics for ALL clusters ...")
		} else {
			fmt.Println("Return statistics for current running clusters ...")
		}

	},
}
