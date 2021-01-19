package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	wCfg = viper.New()
)

type ConfigError struct {
	Err     error
	Message string
}

func (r *ConfigError) Error() string {
	return r.Err.Error()
}

func (r *ConfigError) AlreadyExists() bool {
	_, ok := r.Err.(viper.ConfigFileAlreadyExistsError)
	return ok
}

func (r *ConfigError) NotFound() bool {
	_, ok := r.Err.(viper.ConfigFileNotFoundError)
	return ok
}

type WorkspaceConfig struct {
	Name        string
	Description string
	AppsDir     string
}

func setWorkspaceDefaults() {
	wCfg.SetConfigName("inspr.workspace")
	wCfg.SetConfigType("yaml")

	wCfg.SetDefault("AppsDir", AppsDir())
	wCfg.SetDefault("Description", "Add your Workspace description")
}

func AppsDir() string {
	if name := wCfg.GetString("AppsDir"); name != "" {
		return name
	}
	return "apps"
}

func CreateWorkspace(name string) error {
	wCfg.Set("Name", name)

	if err := wCfg.SafeWriteConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to write Workspace config",
		}
	}

	fmt.Printf("Created new Workspace in: %s \n Settings: %+v", wCfg.ConfigFileUsed(), wCfg.AllSettings())
	return nil
}

func InitWorkspace(path string) (*WorkspaceConfig, *ConfigError) {
	var conf WorkspaceConfig

	setWorkspaceDefaults()
	if path != "" {
		wCfg.AddConfigPath(path)
	} else {
		wCfg.AddConfigPath(".")
	}

	if err := wCfg.ReadInConfig(); err != nil {
		return nil, &ConfigError{
			Err:     err,
			Message: "failed to read config",
		}
	}

	if err := wCfg.Unmarshal(&conf); err != nil {
		return nil, &ConfigError{
			Message: "unable to decode into struct",
			Err:     err,
		}
	}

	return &conf, nil
}

func DescribeWorkspace() *ConfigError {
	if wCfg.ConfigFileUsed() == "" {
		return &ConfigError{
			Err:     viper.ConfigFileNotFoundError{},
			Message: "can't describe, workspace config is not located. Use inspr init [name] to create new Workspace",
		}
	}
	fmt.Printf("Workspace config used: %s \n Settings: %+v \n", wCfg.ConfigFileUsed(), wCfg.AllSettings())
	return nil
}
