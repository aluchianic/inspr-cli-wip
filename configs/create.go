package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

func (f *FileRaw) Create(name string) *ConfigError {
	if f.Path == "" {
		panic(fmt.Errorf("can't create config, not path provided"))
	}
	v := viper.New()
	v.SetConfigType("yaml")

	switch f.Definition {

	case "workspace":
		v.AddConfigPath(f.Path)
		v.SetConfigName(name + ".workspace")
		v.SetDefault("AppsDir", "apps")
		v.SetDefault("Description", "Your description goes here")
		v.SetDefault("Applications", []string{})

		if err := createDirs(f.Path); err != nil {
			return &ConfigError{
				Message: "Failed to create directories",
			}
		}
	case "application":
		f.Path = path.Join(f.Path, name)

		v.AddConfigPath(f.Path)
		v.SetConfigName(name + ".application")
		v.SetDefault("Depends", []string{})
		v.SetDefault("Description", "Add your Application description")
		v.SetDefault("Channels", &ChannelYaml{})

		if err := createDirs(f.Path); err != nil {
			return &ConfigError{
				Message: "Failed to create directories",
			}
		}
	}

	_ = v.MergeInConfig()
	if err := v.SafeWriteConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to create Config",
		}
	}
	if err := v.ReadInConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to read config",
		}
	}
	if v.ConfigFileUsed() == "" {
		return &ConfigError{
			Err:     errors.New("config file still not found"),
			Message: "config file still not found",
		}
	}

	f.Path = v.ConfigFileUsed()
	fmt.Printf("Created new '%s' config file: %s \n", f.Definition, f.Path)

	return nil
}

func createDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
