package cmd

import (
	"fmt"
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
	Run: func(cmd *cobra.Command, workspaces []string) {
		if len(workspaces) == 0 {
			fmt.Println("Deploying apps from current directory (workspace)")
		} else {
			fmt.Printf("Deploying apps from workspaces: %v \n", workspaces)
		}
		if len(exclude) > 0 {
			fmt.Printf("Excluding apps from deploying: %v \n", exclude)
		}
	},
}
