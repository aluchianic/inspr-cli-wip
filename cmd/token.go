package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var create bool

func init() {
	rootCmd.AddCommand(tokenCmd)

	tokenCmd.Flags().BoolVarP(&create, "create", "c", false, "Create new token.")
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Create or validate Inpsr token",
	Run: func(cmd *cobra.Command, args []string) {
		if create {
			t := "Hj31-Y762@-LJu1"
			fmt.Printf("Created new token '%s' \n", t)
			viper.Set("Token", t)
		} else {
			fmt.Printf("Your token %s, is valid.", viper.GetString("Token"))
		}
	},
}
