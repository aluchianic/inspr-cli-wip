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

var logger *zap.SugaredLogger

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

// Return `appsDir` value from config
func (w *WorkspaceFiles) getAppsDir() string {
	if !w.Parsed {
		w.Logger.Fatalf("can't retrieve values before parsing, use Parse() method first. \t\"path\": \"%s\", \"parsed\": \"%b\" type: \"%s\"", w.Path, w.Parsed, w.Definition)
	}
	return w.Config.GetString("AppsDir")
}

// Validate and initialize WorkspaceFiles struct
func (w *WorkspaceFiles) init() {
	w.Root = toAbsolute(w.Root)
	if w.ApplicationsFiles == nil {
		w.ApplicationsFiles = ApplicationsFiles{}
	}

}

// Validate and initialize RawConfig struct
func (cfg *RawConfig) init(definition string) {
	// lazy load logger
	if logger == nil {
		logger = logging.Logger()
	}
	cfg.Logger = logger
	cfg.Config = viper.New()

	switch definition {
	case workspace:
		cfg.Definition = definition
	case application:
		cfg.Definition = definition
	default:
		cfg.Logger.Fatalf("Unknown definition for config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
}

// Set RawConfig values
func (cfg *RawConfig) load(configPath string) {
	cfg.Path = configPath

	cfg.Logger.Debugf("Loaded config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

// Adds application to WorkspaceFiles struct
func (w *WorkspaceFiles) addApplication(f RawConfig) {
	name := AppName(f.name())
	w.ApplicationsFiles[name] = f
}

// Returns config name based on filename
func (cfg *RawConfig) name() string {
	return strings.Split(path.Base(cfg.Path), ".")[0]
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
