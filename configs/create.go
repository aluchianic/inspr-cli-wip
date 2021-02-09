package configs

import (
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

	case workspace:
		f.load(path.Join(f.Path, name), f.Definition)

		v.AddConfigPath(f.Path)
		v.SetConfigName(name + "." + workspace)
		v.SetDefault("AppsDir", "apps")
		v.SetDefault("Description", "Your description goes here")
		v.SetDefault("Applications", []string{})

		if err := createDirs(f.Path); err != nil {
			return &ConfigError{
				Message: "Failed to create directories",
			}
		}
	case application:
		f.load(path.Join(f.Path, name), f.Definition)

		v.AddConfigPath(f.Path)
		v.SetConfigName(name + "." + application)
		v.SetDefault("Depends", []string{})
		v.SetDefault("Description", "Add your Application description")
		v.SetDefault("Channels", &ChannelYaml{})

		if err := createDirs(f.Path); err != nil {
			return &ConfigError{
				Message: "Failed to create directories",
			}
		}
	}

	if err := f.parse(); err != nil {
		return err
	}

	fmt.Printf("Created new '%s' config file: %s \n", f.Definition, f.Path)
	return nil
}

func createDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
