package cmd

import (
	//"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCommand)
	// Flags
	//configs.AddPathFlag(deployCommand)
	//configs.AddExcludeFlag(deployCommand)

	//deployCommand.Flags().BoolVarP(&all, "all", "a", false, "add all dApps to execution")
	//deployCommand.Flags().StringVarP(&excludeRegExp, "exclude-reg", "E", "", "exclude resources by RegExp from execution")
	//deployCommand.Flags().StringVarP(&includeRegExp, "include-reg", "I", "", "include resources by RegExp into execution")
	//deployCommand.Flags().StringSliceVarP(&include, "include", "i", []string{}, "include resources into execution")
	//deployCommand.Flags().StringSliceVarP(&exclude, "exclude", "c", []string{}, "exclude resources from execution")
}

var deployCommand = &cobra.Command{
	Use:   "deploy [appName]",
	Short: "[Cluster] Deploy Workspace on cluster if no arguments passed assuming that Workspace is current directory.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, apps []string) {
		//var (
		//	w   *configs.Workspace
		//	a   *configs.Application
		//	err *configs.ConfigError
		//)
		//
		//w = configs.InitWorkspace()
		//a, err = w.InitApplication("test")
		//configs.ShowAndExistIfErrorExists(err)
		//fmt.Printf("Current workspace :: %+v \n", w)
		//fmt.Printf("Current application :: %+v \n", a)
	},
}
