package cmd

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"
)

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file
func initConfig() {
	// try to locate workspace
	cm := config.CM()
	_ = cm.LoadConfigs(cm.Flags.WorkspaceDir)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Main(); err != nil {
		util.Errorf(err.Error())
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
