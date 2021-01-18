package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

var (
	all                                     bool
	exclude, include                        []string
	excludeRegExp, includeRegExp, workspace string
)

func init() {
	rootCmd.AddCommand(deployCommand)

	deployCommand.Flags().StringVarP(&workspace, "workspace", "w", "", "set path to workspace")
	deployCommand.Flags().BoolVarP(&all, "all", "a", false, "add all dApps to execution")
	deployCommand.Flags().StringVarP(&excludeRegExp, "exclude-reg", "E", "", "exclude resources by RegExp from execution")
	deployCommand.Flags().StringVarP(&includeRegExp, "include-reg", "I", "", "include resources by RegExp into execution")
	deployCommand.Flags().StringSliceVarP(&include, "include", "i", []string{}, "include resources into execution")
	deployCommand.Flags().StringSliceVarP(&exclude, "exclude", "c", []string{}, "exclude resources from execution")
}

var deployCommand = &cobra.Command{
	Use:   "deploy [appName]",
	Short: "[Cluster] Deploy Workspace on cluster if no arguments passed assuming that Workspace is current directory.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, apps []string) {
		w := configs.Workspace{Path: workspace}
		a := configs.App{Name: apps[0]}

		if !w.Init() {
			panic("Workspace not located.")
		}
		if !a.Init() {
			panic("App not located.")
		}

		fmt.Println("Deploying from workspace ...")
		w.Describe()
		fmt.Println("Deploying app ... ")
		a.Describe()
	},
}
