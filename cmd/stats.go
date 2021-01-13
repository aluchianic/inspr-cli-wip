package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statsCmd)

	statsCmd.Flags().BoolP("all", "a", false, "Return statistics from all clusters")
	statsCmd.Flags().BoolP("life", "l", false, "(?) Return a list with lifecycle info regards running dApps")
}

var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"list"},
	Short:   "[Cluster] Return statistics from clusters, by default returns statistics only for running clusters",
	Run:     func(cmd *cobra.Command, args []string) {},
}
