package hub

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/pkg/command"
)

// hubFindCommand represents the `hub find` command
var hubFindCommand *cobra.Command

var _ = command.RegisterCommandVar(func() {
	hubFindCommand = &cobra.Command{
		Use:   "find [application name]",
		Short: "Find application in Inspr Hub",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, dAppName []string) {
			name := dAppName[0]
			author := "A. S. Puskin"
			desc := "Returns name of planet named after author."

			fmt.Printf("apps \"%s\" \n author %s \n description %s", name, author, desc)
		},
	}
})

var _ = command.RegisterCommandInit(func() {
	HubCommandAddCommand(hubFindCommand)
})
