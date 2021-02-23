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
	cm *Manager
)

// CM returns the instance to the config manager
func CM() *Manager {
	if cm == nil {
		cm = newConfigManager()
	}

	return cm
}

func newWorkspace() *Workspace {
	return &Workspace{
		RawConfig:    &RawConfig{},
		Applications: make(map[string]*RawConfig),
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
func (cm *Manager) CreateConfig(name string, definition string) {
	cfgPath := cm.Config.createPath(name, definition)
	raw := cm.load(cfgPath, definition)
	raw.create()
}

// Loads workspace and all application config inside `WorkspaceConfig.AppsDir` and 2 level down
func (cm *Manager) LoadConfigs(root string) *ConfigError {
	cm.Config.Root = toAbsolute(root)

	var matches []string
	// search workspace files
	matches, _ = filepath.Glob(path.Join(cm.Config.Root, workspaceFileName))
	if len(matches) == 0 {
		return ErrNotFound(workspace, cm.Config.Root)
	}

	cm.load(matches[0], workspace)

	// search application files
	matches, _ = filepath.Glob(path.Join(cm.Config.Root, "**/**", applicationFileName))
	for _, match := range matches {
		cm.load(match, application)
	}

	return nil
}

// Load RawConfig structure
func (cm *Manager) load(path string, definition string) *RawConfig {
	raw := &RawConfig{
		Path:   path,
		Config: NewConfig(),
	}

	switch definition {
	case workspace:
		raw.Definition = definition

		raw.Config.SetDefault("AppsDir", "apps")
		raw.Config.SetDefault("Description", "Your description goes here")
		raw.Config.SetDefault("Applications", []string{})

		cm.Config.RawConfig = raw
	case application:
		raw.Definition = definition

		raw.Config.SetDefault("Depends", []string{})
		raw.Config.SetDefault("Description", "Add your Application description")
		raw.Config.SetDefault("Channels", &ChannelYaml{})

		cm.Config.addApplication(raw)
	default:
		util.Errorf("Unknown definition for config \t\"path\": \"%s\"\t\"type\": \"%s\"", raw.Path, raw.Definition)
	}

	return raw
}

// Returns Application config in case if exists in Workspace, empty string otherwise
func (w *Workspace) search(name string) *RawConfig {
	for appName, _ := range w.Applications {
		if appName == name {
			rawCfg := w.Applications[appName]
			return rawCfg
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

// Adds application to Workspace struct
func (w *Workspace) addApplication(cfg *RawConfig) {
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
