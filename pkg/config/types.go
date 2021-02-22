package config

import (
	"github.com/spf13/viper"
)

// Workspace contains raw config definition files for Workspace and it Applications
type Workspace struct {
	RawConfig
	Applications map[string]RawConfig
	Root         string
}

// Raw config definition files
type RawConfig struct {
	Path       string
	Parsed     bool
	Definition string
	Config     *viper.Viper
}

/////////// CONSTANTS //////////////////
const (
	homedirFolder       = ".inspr"
	workspace           = "workspace"
	application         = "application"
	workspaceFileName   = "*." + workspace + ".yaml"
	applicationFileName = "*." + application + ".yaml"
)

/////////// Configs ////////////////////
// TODO NOT USED: MainCliYaml is yaml for main cli settings $HOME/.inspr/inspr.config.yaml
type MainCliYaml struct {
	Version string `yaml:"version"`
	Account string `yaml:"account"`
	Token   string `yaml:"token"`
}

// WorkspaceYaml is a yaml for a channel
type WorkspaceYaml struct {
	Description  string   `yaml:"description"`
	AppsDir      string   `yaml:"appsdir"`
	Applications []string `yaml:"applications"`
}

// ApplicationYaml is a yaml for a channel
type ApplicationYaml struct {
	Name         string   `yaml:"name"`
	Id           string   `yaml:"id"`
	Dependencies []string `yaml:"dependencies"`
	ChannelYaml  `yaml:"channels"`
}

// ChannelYaml is a yaml for a channel
type ChannelYaml struct {
	Version   float32   `yaml:"version"`
	Namespace string    `yaml:"namespace"`
	Channels  []Channel `yaml:"channels"`
}

// Channel is the definition for a channel
type Channel struct {
	Name       string `yaml:"name"`
	Avropath   string `yaml:"avropath"`
	Partitions int32  `yaml:"partitions"`
}
