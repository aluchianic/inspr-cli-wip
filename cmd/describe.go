package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/configs"
)

func init() {
	rootCmd.AddCommand(describeCmd)
	hubCmd.AddCommand(describeCmd)
}

var describeCmd = &cobra.Command{
	Use:   "describe [appName]",
	Short: "DescribeApp dApp with it dependencies, channel types and third parties",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := configs.App{Name: args[0]}
		a.Init()
		a.Describe()
	},
}
