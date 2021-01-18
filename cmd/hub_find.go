package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	hubCmd.AddCommand(hubFindCmd)
}

var hubFindCmd = &cobra.Command{
	Use:   "find [dApp name]",
	Short: "Init dApp in Inpsr Hub",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, dAppName []string) {
		name := dAppName[0]
		author := "A. S. Puskin"
		desc := "Returns name of planet named after author."

		fmt.Printf("app '%s' \n author %s \n description %s", name, author, desc)
	},
}
