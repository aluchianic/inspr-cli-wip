package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Sets default values and creates new config file
func (f *RawConfig) Create(name string, relativePath string, definition string) *ConfigError {
	filename := name + "." + definition + ".yaml"

	switch definition {

	case workspace:
		f.load(path.Join(relativePath, filename), definition)

		f.Config.SetDefault("AppsDir", "apps")
		f.Config.SetDefault("Description", "Your description goes here")
		f.Config.SetDefault("Applications", []AppName{})
	case application:
		f.load(path.Join(relativePath, name, filename), definition)

		f.Config.SetDefault("Depends", []string{})
		f.Config.SetDefault("Description", "Add your Application description")
		f.Config.SetDefault("Channels", &ChannelYaml{})
	default:
		return &ConfigError{
			Message: "unknown definition: '" + definition + "'",
		}
	}

	if err := createDirs(f.Path); err != nil {
		return &ConfigError{
			Message: "Failed to create directories for: " + f.Path,
		}
	}

	if err := f.create(); err != nil {
		return err
	}

	fmt.Printf("Created new '%s' config file: %s \n", f.Definition, f.Path)
	return nil
}

// Creates new config file based on its' Path
func (f *RawConfig) create() *ConfigError {
	err := f.Config.MergeInConfig()
	if err = f.Config.SafeWriteConfigAs(f.Path); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to create new `" + f.Definition + "` config, under: " + f.Path,
		}
	}

	return nil
}

// Create directories recursively
func createDirs(path string) error {
	dir, _ := filepath.Split(path)
	return os.MkdirAll(dir, os.ModePerm)
}
