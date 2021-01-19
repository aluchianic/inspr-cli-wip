package configs

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
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

func pathToApp(name string) string {
	return path.Join(AppsDir(), name)
}

// []string AppNames
func InitApp(name string) (*AppConfig, *ConfigError) {
	if name == "" {
		return nil, &ConfigError{
			Err: errors.New("name is required for App"),
		}
	}
	var conf AppConfig

	setAppDefaults(name)
	vApp.AddConfigPath(pathToApp(name))

	if err := vApp.ReadInConfig(); err != nil {
		return nil, &ConfigError{
			Err:     err,
			Message: "failed to read App config",
		}
	}
	if err := wCfg.Unmarshal(&conf); err != nil {
		return nil, &ConfigError{
			Message: "unable to decode App config into struct",
			Err:     err,
		}
	}

	return &conf, nil
}

// strings[] - App names
func CreateApp(name string) *ConfigError {
	err := createDirIfNotExists(AppsDir())
	err = createDirIfNotExists(pathToApp(name))
	if err = vApp.SafeWriteConfig(); err != nil {
		return &ConfigError{
			Message: "failed to write App config",
			Err:     err,
		}
	}
	return nil
}

func DescribeApp() *ConfigError {
	if vApp.ConfigFileUsed() == "" {
		return &ConfigError{
			Err:     viper.ConfigFileNotFoundError{},
			Message: "can't describe, App config is not located. Use inspr `init [workspace] -app`  to create new App",
		}
	}
	fmt.Printf("App config used: %s, \n Settings: %+v\n", vApp.ConfigFileUsed(), vApp.AllSettings())
	return nil
}
