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
		configPath string
		filename   = name + "." + definition + ".yaml"
		logger     = logging.Logger()
	)

	switch definition {
	case workspace:
		configPath = path.Join(relativePath, filename)
	case application:
		configPath = path.Join(relativePath, name, filename)
	}

	f.load(configPath, definition, logger)
	return f.create()
}

// Creates new config file based on its' Path
func (f *RawConfig) create() *ConfigError {
	if err := createDirs(f.Path); err != nil {
		f.Logger.Error("failed to create directories", zap.String("path", f.Path), zap.String("type", f.Definition))
	}

	err := f.Config.MergeInConfig()
	if err = f.Config.SafeWriteConfigAs(f.Path); err != nil {
		f.Logger.Error("failed to create new config", zap.String("path", f.Path), zap.String("type", f.Definition))
	}

	f.Logger.Info("Created new config", zap.String("path", f.Path), zap.String("type", f.Definition))

	return nil
}

// Create directories recursively
func createDirs(path string) error {
	dir, _ := filepath.Split(path)
	return os.MkdirAll(dir, os.ModePerm)
}
