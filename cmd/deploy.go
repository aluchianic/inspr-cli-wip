package cmd

import (
	"github.com/spf13/cobra"
)

var exclude []string

func init() {
	rootCmd.AddCommand(deployCommand)

	deployCommand.Flags().StringSliceVarP(&exclude, "exclude", "e", []string{}, "exclude resources from deploy")
}

var deployCommand = &cobra.Command{
	Use:   "deploy [workspaces...]",
	Short: "[Cluster] Deploy Workspace on cluster if no arguments passed assuming that Workspace is current directory.",
	Args:  cobra.ArbitraryArgs,
	Run:   func(cmd *cobra.Command, args []string) {},
}
