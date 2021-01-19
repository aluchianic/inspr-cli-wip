package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/configs"
	"os"
)

var (
	force bool
	w     string
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Remove dApp without graceful shutdown of service.")
	removeCmd.Flags().StringVarP(&w, "workspace", "w", "", "Set path to workspace")
	// --safe [--default]
}

var removeCmd = &cobra.Command{
	Use:   "remove [appName...]",
	Short: "[Cluster] Remove deployed dApp from cluster",
	Args:  cobra.MinimumNArgs(1),
	Long:  "By default command tries to remove in 'safe' mode, awaiting to close all connections",
	Run: func(cmd *cobra.Command, appNames []string) {
		var (
			aConf *configs.AppConfig
			err   *configs.ConfigError
		)
		_, err = configs.InitWorkspace(workspace)
		aConf, err = configs.InitApp(appNames[0])

		err = configs.DescribeWorkspace()
		err = configs.DescribeApp()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Message)
			os.Exit(1)
		}
		fmt.Printf("Removing app with name: %s", aConf.Name)
	},
}
