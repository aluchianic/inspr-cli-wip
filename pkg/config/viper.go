package config

import (
	"inspr-cli/pkg/util"

	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Returns new Viper instance
func NewConfig() *viper.Viper {
	return viper.New()
}

func (cfg *RawConfig) exists(err error) bool {
	_, ok := err.(viper.ConfigFileAlreadyExistsError)
	return ok
}

// Creates new config file and directories based on its' Path
func (cfg *RawConfig) create() {
	if err := createDirs(cfg.Path); err != nil {
		util.Errorf("failed to create directories \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}
	err := cfg.Config.MergeInConfig()
	if err = cfg.Config.SafeWriteConfigAs(cfg.Path); err != nil {
		if cfg.exists(err) {
			util.Errorf("config already exists \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
		}
		util.Errorf("failed to create new config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
	}

	util.Debugf("Created new config \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)
}

// Sets default values for Configs
func (cfg *RawConfig) setConfigDefaults() {
	util.Debugf("Setting config defaults \t\"path\": \"%s\"\t\"type\": \"%s\"", cfg.Path, cfg.Definition)

	switch cfg.Definition {
	case workspace:
		cfg.Config.SetDefault("AppsDir", "apps")
		cfg.Config.SetDefault("Description", "Your description goes here")
		cfg.Config.SetDefault("Applications", []string{})
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
