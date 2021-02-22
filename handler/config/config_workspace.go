package config

import (
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"
)

// workspaceCommand represents the `config workspace` command
var workspaceCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	workspaceCommand = &cobra.Command{
		Use:   "workspace [workspace name]",
		Args:  cobra.ExactArgs(1),
		Short: "Initialize fresh Inspr workspace config",
		Run: func(_ *cobra.Command, args []string) {
			cm := config.CM()
			err := cm.Load(cm.Flags.WorkspaceDir)
			if err != nil && err.NotFound() {
				cm.Create(args[0], "workspace")
				util.Infof("Created new workspace in: %s", cm.Config.Workspace.Path)
			}
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(workspaceCommand)
})
