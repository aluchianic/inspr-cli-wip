package config

// Parses all WorkspaceFiles
func (w *WorkspaceFiles) Parse() {
	// parse workspace
	w.RawConfig.parse(WorkspaceYaml{})

	// parse application
	for _, fileRaw := range w.ApplicationsFiles {
		fileRaw.parse(ApplicationYaml{})
	}
}

// Parses config file according to interface
func (cfg *RawConfig) parse(i interface{}) {
	cfg.unmarshal(&i)
	cfg.read()

	cfg.Parsed = true
	cfg.Logger.Debugf("Parsed config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

// Unmarshal config
func (cfg *RawConfig) unmarshal(i interface{}) {
	if err := cfg.Config.Unmarshal(&i); err != nil {
		cfg.Logger.Fatalf("failed to unmarshal \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
}

// Read config in memory
func (cfg *RawConfig) read() {
	cfg.Config.SetConfigFile(cfg.Path)

	if err := cfg.Config.MergeInConfig(); err != nil {
		cfg.Logger.Fatalf("failed to merge config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
	if err := cfg.Config.ReadInConfig(); err != nil {
		cfg.Logger.Fatalf("failed to read config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
}
