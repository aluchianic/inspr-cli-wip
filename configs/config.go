package configs

import (
	"inspr-cli/log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Creates config files and folders
func (w *WorkspaceFiles) Create(name string, definition string) {
	var rawConfig = RawConfig{}

	rawConfig.init(definition)
	rawConfig.setConfigDefaults()
	rawConfig.setConfigPath(w.createPath(name, definition))

	if definition == workspace {
		w.RawConfig = rawConfig
	} else {
		w.addApplication(rawConfig)
	}
	rawConfig.create()
}

// Loads workspace and all application configs inside `WorkspaceConfig.AppsDir` and 2 level down
func (w *WorkspaceFiles) Load() *ConfigError {
	// validate and initialize default values for workspace
	w.init()

	var matches []string
	// search workspace files
	matches, _ = filepath.Glob(path.Join(w.Root, workspaceFileName))
	if len(matches) == 0 {
		return ErrNotFound(workspace, w.Root)
	}

	w.RawConfig.init(workspace)
	w.RawConfig.setConfigPath(matches[0])

	// search application files
	matches, _ = filepath.Glob(path.Join(w.Root, "**/**", applicationFileName))
	for _, match := range matches {
		app := RawConfig{}
		app.init(application)
		app.setConfigPath(match)
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
	cfg.defineType(definition)

	cfg.Logger = log.Logger
	cfg.Config = NewConfig()
	cfg.Logger.Debugf("init raw config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

func (cfg *RawConfig) defineType(definition string) {
	switch definition {
	case workspace:
	case application:
	default:
		cfg.Logger.Fatalf("Unknown definition for config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
	cfg.Definition = definition
}

// Set RawConfig values
func (cfg *RawConfig) setConfigPath(configPath string) {
	cfg.Path = configPath

	cfg.Logger.Debugf("set config path \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
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

// Creates new path to file
func (w *WorkspaceFiles) createPath(name string, definition string) string {
	var filename = name + "." + definition + ".yaml"

	switch definition {
	case workspace:
		return path.Join(w.Root, filename)
	case application:
		return path.Join(w.Root, w.getAppsDir(), name, filename)
	default:
		return ""
	}
}
