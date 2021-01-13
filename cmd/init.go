package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inspr-cli/helpers"
	"os"
	"path"
)

var (
	projectName string

	initCmd = &cobra.Command{
		Use:   "init [name]",
		Short: "[Workspace] Initialize Inspr application",
		Long: `Initialize (inspr-cli init) will create a new application
  * If a name is provided, a directory with that name will be created in the current directory;
  * If no name is provided, the current directory will be assumed;
`,
		Run: func(_ *cobra.Command, args []string) {
			fmt.Printf("%v", args)
			projectPath, err := initializeProject(args)
			if err != nil {
				panic(fmt.Errorf("Failed to initialize project: %s ", err))
			}
			fmt.Printf("Your Inspr application is ready at\n%s\n", projectPath)
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func initializeProject(args []string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}

	project := &helpers.Workspace{
		AbsolutePath: wd,
		Name:         path.Base(projectName),
	}

	if err := project.Create(); err != nil {
		return "", err
	}

	return project.AbsolutePath, nil
}
