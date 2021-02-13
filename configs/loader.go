package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"inspr-cli/logging"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Loads workspace and all application configs inside `WorkspaceConfig.AppsDir` and 2 level down
func (w *WorkspaceFiles) Load() *ConfigError {
	var (
		logger  = logging.Logger()
		matches []string
	)
	// set Default
	if w.Root == "" {
		w.Root = toAbsolute("") // cwd
	}
	w.ApplicationsFiles = ApplicationsFiles{}

	// Load Workspace
	matches, _ = filepath.Glob(path.Join(w.Root, workspaceFileName))
	if len(matches) == 0 {
		return ErrNotFound(workspace, w.Root)
	}
	w.load(matches[0], workspace, logger)

	// load Applications
	matches, _ = filepath.Glob(path.Join(w.Root, "**/**", applicationFileName))
	for _, match := range matches {
		name := AppName(strings.Split(path.Base(match), ".")[0])

		app := RawConfig{}
		app.load(match, application, logger)
		app.Logger = logger

		w.ApplicationsFiles[name] = app
	}

	return nil
}

// Set RawConfig values
// TODO: replace configPath \w []string{} - patterns
func (f *RawConfig) load(configPath string, definition string, logger *zap.Logger) {
	f.Path = configPath
	f.Definition = definition
	f.Config = viper.New()
	f.Logger = logger

	switch definition {
	case workspace:
		f.Config.SetDefault("AppsDir", "apps")
		f.Config.SetDefault("Description", "Your description goes here")
		f.Config.SetDefault("Applications", []AppName{})
	case application:
		f.Config.SetDefault("Depends", []string{})
		f.Config.SetDefault("Description", "Add your Application description")
		f.Config.SetDefault("Channels", &ChannelYaml{})
	default:
		f.Logger.Fatal("Unknown definition for config", zap.String("type", f.Definition))
	}

	f.Logger.Info("Loaded config", zap.String("path", f.Path), zap.String("type", f.Definition))
}

// Returns config name based on filename
func (f *RawConfig) name() string {
	return strings.Split(path.Base(f.Path), ".")[0]
}

// return absolute path, wd in case of - ""
func toAbsolute(p string) (abs string) {
	var res string
	if path.IsAbs(p) {
		res = p
	} else {
		dir, err := os.Getwd()
		if err != nil {
			return ""
		}
		res = path.Join(dir, p)
	}

	return res
}
