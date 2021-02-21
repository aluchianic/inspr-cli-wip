package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"os"
)

// appCommand represents the `config app` command
var appCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	appCommand = &cobra.Command{
		Use:   "app [application names...]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Initialize fresh application configs",
		Run: func(_ *cobra.Command, appNames []string) {
			workspace := config.WorkspaceFiles{
				Root: workspacePath,
			}

			err := workspace.Load()
			if err != nil {
				// TODO: !!
				fmt.Errorf("err : %v", err)
				os.Exit(1)
			}

			// Parse workspace allowing read content of WorkspaceConfig to create Application
			workspace.Parse()

			for _, app := range appNames {
				workspace.Create(app, "application")
				workspace.Logger.Infof("Created new application in workspace: %s", workspace.Path)
			}
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(appCommand)
})
