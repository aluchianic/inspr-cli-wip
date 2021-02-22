package config

import (
	"inspr-cli/pkg/util"
)

// Parses all Workspace
func (w *Workspace) Parse() {
	// parse workspace
	w.Workspace.parse(WorkspaceYaml{})

	// parse application
	for _, fileRaw := range w.Applications {
		fileRaw.parse(ApplicationYaml{})
	}
}

// Parses config file according to interface
func (cfg *RawConfig) parse(i interface{}) {
	cfg.unmarshal(&i)
	cfg.read()

	cfg.Parsed = true
	util.Debugf("Parsed config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

// Unmarshal config
func (cfg *RawConfig) unmarshal(i interface{}) {
	if err := cfg.Config.Unmarshal(&i); err != nil {
		util.Errorf("failed to unmarshal \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
}

// Read config in memory
func (cfg *RawConfig) read() {
	cfg.Config.SetConfigFile(cfg.Path)

	if err := cfg.Config.MergeInConfig(); err != nil {
		util.Errorf("failed to merge config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
	if err := cfg.Config.ReadInConfig(); err != nil {
		util.Errorf("failed to read config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
}
