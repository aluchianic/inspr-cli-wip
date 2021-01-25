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
		var (
			appName = args[0]
			w       *configs.Workspace
			a       *configs.Application
			err     *configs.ConfigError
		)

		w = configs.InitWorkspace()
		a, err = w.InitApplication(appName)
		configs.ShowAndExistIfErrorExists(err)
		err = a.Describe()
		configs.ShowAndExistIfErrorExists(err)
	},
}
