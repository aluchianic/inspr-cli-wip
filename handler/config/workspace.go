package config

import (
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"

	"github.com/spf13/cobra"
)

// workspaceCommand represents the `config workspace` command
var workspaceCommand *cobra.Command

func CreateWorkspace(_ *cobra.Command, args []string) {
	cm := config.CM()
	err := cm.LoadConfigs(cm.Flags.WorkspaceDir)

	if err == nil {
		util.Errorf("workspace already exists : %s", cm.Config.Path)
	} else {
		if err.NotFound() {
			cm.CreateConfig(args[0], "workspace")
			util.Infof("Created new workspace in: %s", cm.Config.Path)
		}
	}
}

var _ = command.RegisterCommandVar(func() {
	workspaceCommand = &cobra.Command{
		Use:   "workspace [workspace name]",
		Args:  cobra.ExactArgs(1),
		Short: "Initialize fresh Inspr workspace config",
		Run:   CreateWorkspace,
	}
})

var _ = command.RegisterCommandInit(func() {
	ConfigCommandAddCommand(workspaceCommand)
})
