package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/configs"
	"os"
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
		_, err := configs.InitApp(args[0])
		err = configs.DescribeApp()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	},
}
