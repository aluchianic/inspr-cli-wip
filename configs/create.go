package configs

import (
	"inspr-cli/logging"
	"os"
	"path"
	"path/filepath"
)

// Creates config files and folders
func (w *WorkspaceFiles) Create(name string, definition string) {
	if logger == nil {
		logger = logging.Logger()
	}

	var rawConfig = RawConfig{}

	rawConfig.init(definition)
	rawConfig.setConfigDefaults()
	rawConfig.load(w.createPath(name, definition))

	if definition == workspace {
		w.RawConfig = rawConfig
	} else {
		w.addApplication(rawConfig)
	}
	rawConfig.create()
}

// Creates new config file and directories based on its' Path
func (cfg *RawConfig) create() {
	if err := createDirs(cfg.Path); err != nil {
		cfg.Logger.Fatalf("failed to create directories \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}

	err := cfg.Config.MergeInConfig()
	if err = cfg.Config.SafeWriteConfigAs(cfg.Path); err != nil {
		cfg.Logger.Fatalf("failed to create new config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}

	cfg.Logger.Debugf("Created new config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

// Sets default values for Configs
func (cfg *RawConfig) setConfigDefaults() {
	cfg.Logger.Debugf("Setting config defaults \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)

	switch cfg.Definition {
	case workspace:
		cfg.Config.SetDefault("AppsDir", "apps")
		cfg.Config.SetDefault("Description", "Your description goes here")
		cfg.Config.SetDefault("Applications", []AppName{})
	case application:
		cfg.Config.SetDefault("Depends", []string{})
		cfg.Config.SetDefault("Description", "Add your Application description")
		cfg.Config.SetDefault("Channels", &ChannelYaml{})
	}
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
		return path.Join(w.Root, w.getAppsDir(), name, filename)
	default:
		return ""
	}
}
