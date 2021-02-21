package config

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
)

// workspaceCommand represents the `config workspace` command
var workspaceCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	workspaceCommand = &cobra.Command{
		Use:   "workspace [workspace name]",
		Args:  cobra.ExactArgs(1),
		Short: "Initialize fresh Inspr workspace config",
		Run: func(_ *cobra.Command, args []string) {
			workspace := config.WorkspaceFiles{
				Root: workspacePath,
			}

			if err := workspace.Load(); err != nil && err.NotFound() {
				workspace.Create(args[0], "workspace")
				workspace.Logger.Infof("Created new workspace in: %s", workspace.Path)
			}
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(workspaceCommand)
})
