package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"inspr-cli/pkg/command"
)

var versionCmd *cobra.Command

var _ = command.RegisterCommandVar(func() {
	// rootCommand represents the base command when called without any subcommands
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Inspr cli version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("v%s\n", viper.GetString("version"))
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	RootCommandAddCommand(versionCmd)
})
