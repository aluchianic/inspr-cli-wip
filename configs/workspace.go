package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	wCfg = viper.New()
)

type Workspace struct {
	Name string
	Path string
}

func setWorkspaceDefaults(name string) {
	wCfg.SetConfigType("yaml")
	wCfg.SetDefault("Name", name)
	wCfg.SetConfigName("inspr.workspace")

	wCfg.SetDefault("AppsDir", AppsDir())
	wCfg.SetDefault("Token", viper.GetString("Token"))
	wCfg.SetDefault("Description", "Add your Workspace description")
}

func AppsDir() string {
	if name := wCfg.GetString("AppsDir"); name != "" {
		return name
	}
	return "apps"
}

func (w *Workspace) Create() {
	if err := wCfg.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			fmt.Println("Workspace already initialized")
		} else {
			panic(fmt.Errorf("Failed to write config file: %s \n", err))
		}
	}
	fmt.Printf("Created new Workspace in: %s \n Name: %s \n Description: %s \n AppsDir: %s", wCfg.ConfigFileUsed(), wCfg.GetString("Name"), wCfg.GetString("Description"), AppsDir())
}

func (w *Workspace) Init() bool {
	setWorkspaceDefaults(w.Name)
	if w.Path != "" {
		wCfg.AddConfigPath(w.Path)
	} else {
		wCfg.AddConfigPath(".")
	}

	if err := wCfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}
	fmt.Printf("Workspace config file : %s \n", wCfg.ConfigFileUsed())

	return true
}
func (w *Workspace) Describe() {
	if wCfg.ConfigFileUsed() == "" {
		fmt.Println("Can't describe, workspace config is not located. Use inspr init [name] to create new Workspace")
		return
	}
	fmt.Printf("Workspace name: %s \n Apps Dir: %s \n", wCfg.GetString("Name"), AppsDir())
}
