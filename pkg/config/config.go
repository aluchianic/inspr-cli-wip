package config

import (
	"inspr-cli/pkg/util"

	"os"
	"path"
	"path/filepath"
	"strings"
)

type Manager struct {
	Config *Workspace
	Flags  *ConfigFlags
}

var (
	gcm *Manager
)

// CM returns the instance to the config manager
func CM() *Manager {
	if gcm == nil {
		gcm = newConfigManager()
	}

	return gcm
}

func newWorkspace() *Workspace {
	return &Workspace{
		RawConfig:    RawConfig{},
		Applications: make(map[string]RawConfig),
		Root:         "",
	}
}

func newConfigManager() *Manager {
	configManager := &Manager{
		Config: newWorkspace(),
		Flags:  newWorkspaceFlags(),
	}

	return configManager
}

// Creates config files and folders
func (gcm *Manager) Create(name string, definition string) {
	cfg := gcm.Config
	cfgPath := cfg.createPath(name, definition)

	rawConfig := RawConfig{}
	rawConfig.init(definition)
	rawConfig.setConfigDefaults()
	rawConfig.setConfigPath(cfgPath)

	switch definition {
	case workspace:
		cfg.RawConfig = rawConfig
	case application:
		cfg.addApplication(rawConfig)
	}

	rawConfig.create()
}

// Loads workspace and all application config inside `WorkspaceConfig.AppsDir` and 2 level down
func (gcm *Manager) Load(root string) *Error {
	cfg := gcm.Config
	cfg.Root = toAbsolute(root)

	var matches []string
	// search workspace files
	matches, _ = filepath.Glob(path.Join(cfg.Root, workspaceFileName))
	if len(matches) == 0 {
		return ErrNotFound(workspace, cfg.Root)
	}

	cfg.init(workspace)
	cfg.setConfigPath(matches[0])

	// search application files
	matches, _ = filepath.Glob(path.Join(cfg.Root, "**/**", applicationFileName))
	for _, match := range matches {
		app := RawConfig{}
		app.init(application)
		app.setConfigPath(match)

		cfg.addApplication(app)
	}

	return nil
}

// Returns Application config in case if exists in Workspace, empty string otherwise
func (w *Workspace) search(name string) *RawConfig {
	for appName, _ := range w.Applications {
		if appName == name {
			rawCfg := w.Applications[appName]
			return &rawCfg
		}
	}

	return nil
}

// Return `appsDir` value from config
func (w *Workspace) getAppsDir() string {
	if !w.Parsed {
		util.Errorf("can't retrieve values before parsing, use Parse() method first. \t\"path\": \"%s\", \"parsed\": \"%b\" type: \"%s\"", w.Path, w.Parsed, w.Definition)
	}
	return w.Config.GetString("AppsDir")
}

// Validate and initialize RawConfig struct
func (cfg *RawConfig) init(definition string) {
	cfg.defineType(definition)

	cfg.Config = NewConfig()
	util.Debugf("init raw config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

func (cfg *RawConfig) defineType(definition string) {
	switch definition {
	case workspace:
	case application:
	default:
		util.Errorf("Unknown definition for config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
	cfg.Definition = definition
}

// Set RawConfig values
func (cfg *RawConfig) setConfigPath(configPath string) {
	cfg.Path = configPath

	util.Debugf("set config path \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

// Adds application to Workspace struct
func (w *Workspace) addApplication(cfg RawConfig) {
	w.Applications[cfg.name()] = cfg
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
func (w *Workspace) createPath(name string, definition string) string {
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
