package cmd

import (
	"github.com/spf13/cobra"
	c "inspr-cli/configs"
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
		c.SetAppName(args[0])
		c.LoadAppConfig()
		c.DescribeApp()
	},
}
