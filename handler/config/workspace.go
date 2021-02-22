package config

import (
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"

	"github.com/spf13/cobra"
)

// workspaceCommand represents the `config workspace` command
var workspaceCommand *cobra.Command

func CreateWorkspace(_ *cobra.Command, args []string) error {
	cm := config.CM()

	err := cm.Load(cm.Flags.WorkspaceDir)
	if err != nil && err.NotFound() {
		cm.Create(args[0], "workspace")
		util.Infof("Created new workspace in: %s", cm.Config.Path)
	}

	return nil
}

var _ = command.RegisterCommandVar(func() {
	workspaceCommand = &cobra.Command{
		Use:   "workspace [workspace name]",
		Args:  cobra.ExactArgs(1),
		Short: "Initialize fresh Inspr workspace config",
		RunE:  CreateWorkspace,
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(workspaceCommand)
})
