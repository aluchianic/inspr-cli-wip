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

var logger *zap.Logger

// Loads workspace and all application configs inside `WorkspaceConfig.AppsDir` and 2 level down
func (w *WorkspaceFiles) Load() *ConfigError {
	// validate and initialize default values for workspace
	w.init()

	// search workspace files
	var matches []string
	matches, _ = filepath.Glob(path.Join(w.Root, workspaceFileName))
	if len(matches) == 0 {
		return ErrNotFound(workspace, w.Root)
	}

	w.RawConfig.init(workspace)
	w.load(matches[0])

	// search application files
	matches, _ = filepath.Glob(path.Join(w.Root, "**/**", applicationFileName))
	for _, match := range matches {
		app := RawConfig{}
		app.init(application)
		app.load(match)
		w.addApplication(app)
	}

	return nil
}

// Returns Application config in case if exists in Workspace, empty string otherwise
func (w *WorkspaceFiles) search(name string) *RawConfig {
	n := AppName(name)

	for appName, _ := range w.ApplicationsFiles {
		if appName == n {
			rawCfg := w.ApplicationsFiles[n]
			return &rawCfg
		}
	}

	return nil
}

// Validate and initialize WorkspaceFiles struct
func (w *WorkspaceFiles) init() {
	w.Root = toAbsolute(w.Root)
	if w.ApplicationsFiles == nil {
		w.ApplicationsFiles = ApplicationsFiles{}
	}

}

// Validate and initialize RawConfig struct
func (f *RawConfig) init(definition string) {
	// lazy load logger
	if logger == nil {
		logger = logging.Logger()
		logger.Info("lazy load logger")
	}

	f.Config = viper.New()
	f.Logger = logger

	switch definition {
	case workspace:
		f.Definition = definition
	case application:
		f.Definition = definition
	default:
		f.Logger.Fatal("Unknown definition for config", zap.String("type", definition))
	}

	f.setConfigDefaults()
}

// Set RawConfig values
func (f *RawConfig) load(configPath string) {
	f.Path = configPath

	f.Logger.Info("Loaded config", zap.String("path", f.Path), zap.String("type", f.Definition))
}

// Sets default values for Configs
func (f *RawConfig) setConfigDefaults() {
	f.Logger.Info("Setting config defaults", zap.String("type", f.Definition))
	switch f.Definition {
	case workspace:
		f.Config.SetDefault("AppsDir", "apps")
		f.Config.SetDefault("Description", "Your description goes here")
		f.Config.SetDefault("Applications", []AppName{})
	case application:
		f.Config.SetDefault("Depends", []string{})
		f.Config.SetDefault("Description", "Add your Application description")
		f.Config.SetDefault("Channels", &ChannelYaml{})
	}
}

// Adds application to WorkspaceFiles struct
func (w *WorkspaceFiles) addApplication(f RawConfig) {
	name := AppName(f.name())
	w.ApplicationsFiles[name] = f
}

// Returns config name based on filename
func (f *RawConfig) name() string {
	return strings.Split(path.Base(f.Path), ".")[0]
}

// return absolute path, pwd in case if arg is empty string
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
