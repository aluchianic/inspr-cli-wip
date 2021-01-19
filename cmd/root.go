package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/configs"
	"os"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "inspr",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
inspr ... .`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("EXECUTE ::", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(configs.InitCLi())
}
