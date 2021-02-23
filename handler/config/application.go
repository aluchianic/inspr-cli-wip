package config

import (
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"

	"github.com/spf13/cobra"
)

// appCommand represents the `config app` command
var appCommand *cobra.Command

func CreateApp(_ *cobra.Command, appNames []string) error {
	cm := config.CM()

	err := cm.LoadConfigs(cm.Flags.WorkspaceDir)
	if err != nil {
		return err
	}

	// Parse workspace allowing read content of WorkspaceConfig to create Application
	cm.Config.Parse()
	for _, app := range appNames {
		cm.CreateConfig(app, "application")
		util.Infof("Created new application in workspace: %s", cm.Config.Path)
	}
	return nil
}

var _ = command.RegisterCommandVar(func() {
	appCommand = &cobra.Command{
		Use:   "app [application names...]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Initialize fresh application configs",
		RunE:  CreateApp,
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(appCommand)
})
