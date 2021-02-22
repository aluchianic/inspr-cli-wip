package cluster

import (
	"inspr-cli/pkg/command"
	"inspr-cli/pkg/config"
	"inspr-cli/pkg/util"

	"github.com/spf13/cobra"
)

// deployCommand represents the `cluster deploy` command
var deployCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	deployCommand = &cobra.Command{
		Use:   "deploy [application name]",
		Short: "Deploy application on cluster",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, apps []string) {
			cm := config.CM()
			err := cm.Load(cm.Flags.WorkspaceDir)
			if err != nil {
				util.Errorf(err.Message)
			}
			cm.Config.Parse()
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	// Flags
	//config.AddPathFlag(deployCommand)
	//config.AddExcludeFlag(deployCommand)

	//deployCommand.Flags().BoolVarP(&all, "all", "a", false, "add all dApps to execution")
	//deployCommand.Flags().StringVarP(&excludeRegExp, "exclude-reg", "E", "", "exclude resources by RegExp from execution")
	//deployCommand.Flags().StringVarP(&includeRegExp, "include-reg", "I", "", "include resources by RegExp into execution")
	//deployCommand.Flags().StringSliceVarP(&include, "include", "i", []string{}, "include resources into execution")
	//deployCommand.Flags().StringSliceVarP(&exclude, "exclude", "c", []string{}, "exclude resources from execution")
	ClusterAddCommand(deployCommand)
})
