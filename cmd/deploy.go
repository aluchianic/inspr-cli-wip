package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	all           bool
	exclude       []string
	excludeRegExp string
	include       []string
	includeRegExp string
)

func init() {
	rootCmd.AddCommand(deployCommand)

	deployCommand.Flags().BoolVarP(&all, "all", "a", false, "add all dApps to execution")
	deployCommand.Flags().StringSliceVarP(&exclude, "exclude", "e", []string{}, "exclude resources from execution")
	deployCommand.Flags().StringVarP(&excludeRegExp, "exclude-reg", "E", "", "exclude resources by RegExp from execution")
	deployCommand.Flags().StringSliceVarP(&include, "include", "i", []string{}, "include resources into execution")
	deployCommand.Flags().StringVarP(&includeRegExp, "include-reg", "I", "", "include resources by RegExp into execution")

	// -w [custom workspace] ??
	// [--tag] -t  			 string
}

var deployCommand = &cobra.Command{
	Use:   "deploy [appNamesFromWorkspace...]",
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
