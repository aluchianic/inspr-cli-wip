package configs

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	workspace  string
	vWorkspace = viper.New()
)

func isPath(s string) bool {
	return strings.Contains(s, string(os.PathSeparator))
}

func getWorkspaceName() string {
	if name := vWorkspace.GetString("Workspace"); name != "" {
		return name
	}
	if workspace != "" && !isPath(workspace) {
		return workspace
	}
	return "default"
}

func setWorkspaceDefaults() {
	name := getWorkspaceName()
	vWorkspace.SetConfigType("yaml")
	vWorkspace.SetDefault("Name", name)
	vWorkspace.SetConfigName("inspr.workspace")

	vWorkspace.SetDefault("AppsDir", GetAppsDir())
	vWorkspace.SetDefault("Token", viper.GetString("Token"))
	vWorkspace.SetDefault("Description", "Add your Workspace description")
}

func AddWorkspaceFlag(c *cobra.Command) {
	c.Flags().StringVarP(&workspace, "workspace", "w", "", "Path to Inpsr Workspace, by default looks into current working directory.")
}

func GetAppsDir() string {
	if name := vWorkspace.GetString("AppsDir"); name != "" {
		return name
	}
	return "apps"
}

func SetWorkspaceName(name string) {
	vWorkspace.Set("Workspace", name)
}

func DescribeWorkspace() {
	if vWorkspace.ConfigFileUsed() != "" {
		fmt.Printf("Workspace name: %s \n Apps Dir: %s \n", getWorkspaceName(), GetAppsDir())
	}
}

func LoadWorkspaceConfig() bool {
	setWorkspaceDefaults()
	if isPath(workspace) {
		vWorkspace.AddConfigPath(workspace)
	} else {
		vWorkspace.AddConfigPath(".")
	}

	if err := vWorkspace.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}
	fmt.Printf("!!!. %s \n", vWorkspace.ConfigFileUsed())

	return true
}

func WriteWorkspaceConfig() {
	if err := vWorkspace.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			fmt.Printf("File already exists at: %s \n", vWorkspace.ConfigFileUsed())
		} else {
			panic(fmt.Errorf("Failed to write config file: %s \n", err))
		}
	}
	fmt.Printf("Created new Workspace in: %s \n Name: %s \n Description: %s \n AppsDir: %s", vWorkspace.ConfigFileUsed(), vWorkspace.GetString("Name"), vWorkspace.GetString("Description"), vWorkspace.GetString("AppsDir"))
}
