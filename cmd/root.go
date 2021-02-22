package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

var rootCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	// rootCommand represents the base command when called without any subcommands
	rootCommand = &cobra.Command{
		Use:   "inspr",
		Short: "Inspr orchestrator",
	}
})

var _ = command.RegisterCommandInit(func() {
	// TODO: add custom usage and help
	// TODO: add initialization to root.PersistentPreRunE (e.g. Logger)
})

func RootCommandAddCommand(c *cobra.Command) {
	rootCommand.AddCommand(c)
}
