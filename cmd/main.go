package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
	"os"
)

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file
func initConfig() {
	// TODO -- change to use CM()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Main(); err != nil {
		//TODO: show error
		os.Exit(1)
	}
}

// Main starts the Inspr cli
// Any initialization to Inspr should be added to root.PersistentPreRunE
func Main() error {
	// Setup commands
	command.Setup()

	// Execute pxc
	return rootCommand.Execute()
}
