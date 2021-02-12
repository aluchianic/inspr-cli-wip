package configs

import (
	"go.uber.org/zap"
	"inspr-cli/logging"
	"os"
	"path"
	"path/filepath"
)

// Sets default values and creates new config file
func (f *RawConfig) Create(name string, relativePath string, definition string) *ConfigError {
	var (
		filename = name + "." + definition + ".yaml"
		logger   = logging.Logger()
	)

	switch definition {

	case workspace:
		f.load(path.Join(relativePath, filename), definition, logger)

		f.Config.SetDefault("AppsDir", "apps")
		f.Config.SetDefault("Description", "Your description goes here")
		f.Config.SetDefault("Applications", []AppName{})
	case application:
		f.load(path.Join(relativePath, name, filename), definition, logger)

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

	f.Logger.Info("Created new config", zap.String("path", f.Path), zap.String("type", f.Definition))
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
