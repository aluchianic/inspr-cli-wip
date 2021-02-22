package config

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"
)

// appCommand represents the `config app` command
var appCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	appCommand = &cobra.Command{
		Use:   "app [application names...]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Initialize fresh application configs",
		Run: func(_ *cobra.Command, appNames []string) {
			cm := config.CM()

			err := cm.Load(cm.Flags.WorkspaceDir)
			if err != nil {
				util.Errorf(err.Message)
			}

			// Parse workspace allowing read content of WorkspaceConfig to create Application
			cm.Config.Parse()
			for _, app := range appNames {
				cm.Create(app, "application")
				util.Infof("Created new application in workspace: %s", cm.Config.Path)
			}
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(appCommand)
})
