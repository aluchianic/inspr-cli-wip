package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Inspr cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", viper.GetString("version"))
	},
}
