package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var vApp = viper.New()

// .app.inspr
type Channels struct {
	In  []string
	Out []string
}

type DApp struct {
	Name      string
	DependsOn []string
	Channels
}

func getAppName() string {
	if vApp.GetString("App") != "" {
		return vApp.GetString("App")
	}
	panic("No app name !")
}

func SetAppName(name string) {
	vApp.Set("App", name)
}

func DescribeApp() {
	if vApp.ConfigFileUsed() != "" {
		fmt.Printf("Config file used: %s \n App: %s \n Description: %s \n", vApp.ConfigFileUsed(), vApp.GetString("App"), vApp.GetString("Description"))
	} else {
		fmt.Printf("Can't resolve App")
	}

}

func setAppDefaults() {
	vApp.SetConfigType("yaml")
	vApp.SetConfigName("inspr.app")
	vApp.SetDefault("Depends", []string{})
	vApp.SetDefault("Channels", &Channels{})
	vApp.SetDefault("Description", "Add your Application description")
	vApp.AddConfigPath(GetAppsDir())
}

func LoadAppConfig() bool {
	if err := createDirIfNotExists(GetAppsDir()); err != nil {
		fmt.Printf("LoadAppConfig. Failed to create Apps Directory")
		return false
	}
	getAppName()
	setAppDefaults()
	if err := vApp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}

	fmt.Printf("!!!. %s \n", vApp.ConfigFileUsed())

	return true
}

func WriteAppConfigToDisk() {
	if err := vApp.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			fmt.Printf("can't create App over existing")
		} else {
			fmt.Printf("failed to write config file: %s \n", err)
		}
	}
}
