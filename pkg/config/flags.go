package config

import "github.com/spf13/pflag"

type Flags struct {
	WorkspaceDir    string
	ApplicationsDir string
}

func (c *Flags) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&c.WorkspaceDir, "workspace", "", "Inspr workspace config path")
	flags.StringVar(&c.ApplicationsDir, "apps", "", "Directory to search applications in")
}
