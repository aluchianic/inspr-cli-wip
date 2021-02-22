package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"log"
	"os"
	"path"
)

type ConfigFlags struct {
	HomeDir         string
	WorkspaceDir    string
	ApplicationsDir string
}

func newWorkspaceFlags() *ConfigFlags {
	// If the cfgFile has not been setup in the arguments, then
	// read it from the HOME directory

	// Check env
	configDir := os.Getenv("INSPRHOMEDIR")
	if len(configDir) == 0 {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("unable to determine home directory: %v\n", err)
		}
		configDir = path.Join(home, homedirFolder)
	}

	return &ConfigFlags{
		HomeDir: configDir,
	}
}

func (c *ConfigFlags) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&c.WorkspaceDir, "workspace", "", "Inspr workspace config path")
	flags.StringVar(&c.ApplicationsDir, "apps", "", "Directory to search applications in")
}
