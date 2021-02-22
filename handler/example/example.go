package example

import (
	"inspr-cli/cmd"
	"inspr-cli/pkg/command"

	"fmt"
	"github.com/spf13/cobra"
)

// exampleCmd represents the example command
var exampleCmd *cobra.Command

var _ = command.RegisterCommandVar(func() {
	exampleCmd = &cobra.Command{
		Use:   "example",
		Short: "This is an example",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Please see inspr example --help for more commands\n")
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	cmd.RootCommandAddCommand(exampleCmd)
})
