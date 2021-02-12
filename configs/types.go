package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// WorkspaceFiles contains raw configs definition files for Workspace and it Applications
type WorkspaceFiles struct {
	RawConfig
	ApplicationsFiles
	Root string
}

type ApplicationsFiles map[AppName]RawConfig

// Raw config definition files
type RawConfig struct {
	Path       string
	Content    []byte
	Parsed     bool
	Definition string
	Config     *viper.Viper
	Logger     *zap.Logger
}

/////////// CONSTANTS //////////////////
const (
	workspace           = "workspace"
	application         = "application"
	workspaceFileName   = "*." + workspace + ".yaml"
	applicationFileName = "*." + application + ".yaml"
)

/////////// Configs ////////////////////
// ??: MainCliYaml is yaml for main cli settings $HOME/.inspr/inspr.config.yaml
type MainCliYaml struct {
	Version string `yaml:"version"`
	Account string `yaml:"account"`
	Token   string `yaml:"token"`
}

// WorkspaceYaml is a yaml for a channel
type WorkspaceYaml struct {
	Description  string    `yaml:"description"`
	AppsDir      string    `yaml:"appsdir"`
	Applications []AppName `yaml:"applications"`
}

// ApplicationYaml is a yaml for a channel
type ApplicationYaml struct {
	Name         AppName  `yaml:"name"`
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

// Application name - unique for workspace
type AppName string

// Channel is the definition for a channel
type Channel struct {
	Name       string `yaml:"name"`
	Avropath   string `yaml:"avropath"`
	Partitions int32  `yaml:"partitions"`
}

// Workspace flag to change path to config file
type WorkspaceFlag = string
