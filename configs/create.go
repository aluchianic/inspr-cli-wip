package configs

import (
	"go.uber.org/zap"
	"inspr-cli/logging"
	"os"
	"path"
	"path/filepath"
)

// Creates config files and folders
func (w *WorkspaceFiles) Create(name string, definition string) *ConfigError {
	if logger == nil {
		logger = logging.Logger()
	}

	var rawConfig = RawConfig{}

	rawConfig.init(definition)
	rawConfig.load(w.createPath(name, definition))

	if definition == workspace {
		w.RawConfig = rawConfig
	} else {
		w.addApplication(rawConfig)
	}

	return rawConfig.create()
}

// Creates new config file and directories based on its' Path
func (f *RawConfig) create() *ConfigError {
	if err := createDirs(f.Path); err != nil {
		f.Logger.Fatal("failed to create directories", zap.String("path", f.Path), zap.String("type", f.Definition))
	}

	err := f.Config.MergeInConfig()
	if err = f.Config.SafeWriteConfigAs(f.Path); err != nil {
		f.Logger.Fatal("failed to create new config", zap.String("path", f.Path), zap.String("type", f.Definition))
	}

	f.Logger.Info("Created new config", zap.String("path", f.Path), zap.String("type", f.Definition))

	return nil
}

// Create directories recursively
func createDirs(path string) error {
	dir, _ := filepath.Split(path)
	return os.MkdirAll(dir, os.ModePerm)
}

// Creates new path to file
func (w *WorkspaceFiles) createPath(name string, definition string) string {
	var filename = name + "." + definition + ".yaml"

	switch definition {
	case workspace:
		return path.Join(w.Root, filename)
	case application:
		return path.Join(w.Root, w.Config.GetString("AppsDir"), name, filename)
	default:
		return ""
	}
}
