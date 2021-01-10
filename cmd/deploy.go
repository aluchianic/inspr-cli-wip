package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	mode    = viper.GetString("Mode")
	exclude []string
)

func init() {
	rootCmd.AddCommand(deployCommand)

	deployCommand.Flags().StringVarP(&mode, "mode", "m", "production", "deploy mode")
	deployCommand.Flags().StringSliceVarP(&exclude, "exclude", "e", []string{}, "exclude resources from deploy")
}

func filter(ss []string, test string) (ret []string) {
	for _, s := range ss {
		if test != s {
			ret = append(ret, s)
		}
	}
	return
}

var deployCommand = &cobra.Command{
	Use:   "deploy [targets for deploy]",
	Short: "Deploy target it distributed cloud storage.",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(exclude) > 0 {
			for _, exc := range exclude {
				args = filter(args, exc)
			}
		}

		fmt.Printf("Deploying with credentials from acc: %s \n", viper.GetString("Acc"))
		fmt.Printf("Deploying mode: %s \n", mode)
		fmt.Printf("Excluded: %v \n", exclude)
		fmt.Printf("Targets: %v \n", args)
	},
}
