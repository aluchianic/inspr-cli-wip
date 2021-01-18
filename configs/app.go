package configs

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"path"
)

var vApp = viper.New()

type App struct {
	Name string
	Id   string
}

type Channels struct {
	In  []string
	Out []string
}

type AppConfig struct {
	Name      string
	Id        string
	DependsOn []string
	Channels
}

func createHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func setAppDefaults(name string) {
	vApp.SetConfigType("yaml")
	vApp.SetConfigName("inspr.app")

	vApp.SetDefault("Depends", []string{})
	vApp.SetDefault("Name", name)
	vApp.SetDefault("Id", createHash(name))
	vApp.SetDefault("Channels", &Channels{})
	vApp.SetDefault("Description", "Add your Application description")
}

func (a *App) path() string {
	return path.Join(AppsDir(), a.Name)
}

func (a *App) Init() bool {
	if a.Name == "" {
		panic("App name is required.")
	}

	setAppDefaults(a.Name)
	vApp.AddConfigPath(a.path())

	if err := vApp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}
	fmt.Printf("Application Config used:  %s \n", vApp.ConfigFileUsed())

	return true
}

func (a *App) Create() {
	err := createDirIfNotExists(AppsDir())
	err = createDirIfNotExists(a.path())
	if err != nil {
		panic("Failed to create App Directories")
	}

	if err := vApp.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			fmt.Println("Trying to create app over existing one.")
		} else {
			panic(fmt.Errorf("failed to write config file: %s \n", err))
		}
	}
}

func (a *App) Describe() {
	if vApp.ConfigFileUsed() != "" {
		fmt.Printf("Config file used: %s \n App: %s \n Description: %s \n", vApp.ConfigFileUsed(), vApp.GetString("Name"), vApp.GetString("Description"))
	} else {
		fmt.Printf("Can't resolve App. Use inspr init -a [name] to init new Application")
	}
}
