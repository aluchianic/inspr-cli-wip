package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Returns new Viper instance
func NewConfig() *viper.Viper {
	return viper.New()
}

// TODO: implement
func (cfg *RawConfig) exists() bool {
	//err := cfg.Config.ReadInConfig()
	//_, ok := err.(viper.ConfigFileAlreadyExistsError)
	//return ok
	return false
}

// Creates new config file and directories based on its' Path
func (cfg *RawConfig) create() {
	if cfg.exists() {
		cfg.Logger.Fatalf("config already exists \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}

	if err := createDirs(cfg.Path); err != nil {
		cfg.Logger.Fatalf("failed to create directories \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
	// TODO: add alreadyExist Error
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
