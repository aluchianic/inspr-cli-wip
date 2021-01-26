package configs

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	wCfg          = viper.New()
	workspaceConf Workspace
	// flags
	workspaceFlag string
	excludeFlag   string
)

// Application name used to identify Application
type AppName string

// Workspace config
type Workspace struct {
	Name         string
	Description  string
	AppsDir      string
	Applications []AppName
}

// Initialize silently, setting default values in case workspace is not located
func InitWorkspace() *Workspace {
	setWorkspaceDefaults()
	if err := locateWorkspaceConfig(); err != nil && !err.NotFound() {
		fmt.Printf("InitWorkspace Handle Error:  %s", err)
		os.Exit(0)
	}
	return &workspaceConf
}

// Creates new workspace in current working dir or in workspaceFlag path
func NewWorkspace(name string) *Workspace {
	wCfg.Set("Name", name)
	wCfg.SetConfigName(name + ".workspace")
	wCfg.SetConfigType("yaml")

	if workspaceFlag != "" {
		wCfg.AddConfigPath(workspaceFlag)
	} else {
		wCfg.AddConfigPath(".")
	}

	w := InitWorkspace()
	if err := w.CreateNewConfig(); err != nil && err.AlreadyExists() {
		fmt.Printf("Workspace already exists in: %s \n", wCfg.ConfigFileUsed())
		return w
	} else {
		ShowAndExistIfErrorExists(err)
	}

	fmt.Printf("Created new Workspace in: %s \n", w.Root())

	return w
}

// Returns workspace root path
func (w *Workspace) Root() string {
	p := getWorkspacePath()
	return strings.Replace(p, filepath.Base(p), "", -1)
}

// Writes in config values from current memory state
func (w *Workspace) WriteInConfig() *ConfigError {
	err := wCfg.MergeInConfig()
	if err = wCfg.WriteConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to write into Workspace config",
		}
	}
	return nil
}

// Creates new config file
func (w *Workspace) CreateNewConfig() *ConfigError {
	err := wCfg.MergeInConfig()
	if err = wCfg.SafeWriteConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to create new Workspace config",
		}
	}
	return nil
}

// Returns if Application exists in workspace config
func (w *Workspace) AppExists(name AppName) bool {
	var found = false
	for _, app := range w.Applications {
		if app == name {
			found = true
		}
	}
	return found
}

// Adds Application to Workspace config
func (w *Workspace) AddApplication(name AppName) *ConfigError {
	w.Applications = append(w.Applications, name)
	wCfg.Set("Applications", w.Applications)
	return w.WriteInConfig()
}

// Returns all settings of Workspace
func (w *Workspace) Describe() *ConfigError {
	if wCfg.ConfigFileUsed() == "" {
		return &ConfigError{
			Err:     viper.ConfigFileNotFoundError{},
			Message: "can't describe, workspace config is not located. Use inspr init [name] to create new Workspace",
		}
	}
	fmt.Printf("Workspace config used: %s \n", wCfg.ConfigFileUsed())
	return nil
}

// Flags to change execution behavior
// Flag to change path to workspace
func AddPathFlag(command *cobra.Command) {
	command.Flags().StringVarP(&workspaceFlag, "workspace-path", "w", "", "set path to workspace")
}

// Flag to exclude some Applications from execution
func AddExcludeFlag(command *cobra.Command) {
	command.Flags().StringVarP(&excludeFlag, "exclude", "e", "", "exclude an app from workspace for executions")
	if excludeFlag != "" {
		fmt.Printf("Removing %s app from execurion : %+v ", excludeFlag, workspaceConf.Applications)
	}
}

// ----------------------------------
// Returns current used Workspace path
func getWorkspacePath() string {
	// If config is already initialized return it path
	if p := wCfg.ConfigFileUsed(); p != "" {
		return p
	}

	var workspaceDir string
	if workspaceFlag != "" {
		workspaceDir = workspaceFlag
	} else if p := wCfg.GetString("WorkspaceDir"); p != "" {
		workspaceDir = p
	}

	p, err := findFile(toAbsolute(workspaceDir), []string{"/*.workspace.yaml"})
	if err != nil {
		return ""
		// TODO:
		//return &ConfigError{
		//	Err:     viper.ConfigFileNotFoundError{},
		//	Message: "failed to locate workspace config",
		//}
	}
	return p
}

// Locate Workspace config
func locateWorkspaceConfig() *ConfigError {
	p := getWorkspacePath()
	wCfg.SetConfigFile(p)

	if err := wCfg.Unmarshal(&workspaceConf); err != nil {
		return &ConfigError{
			Message: "unable to decode into struct",
			Err:     err,
		}
	}

	if err := wCfg.ReadInConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to read config",
		}
	}
	return nil
}

// Sets default values for Workspace config
func setWorkspaceDefaults() {
	wCfg.SetConfigType("yaml")
	wCfg.SetDefault("AppsDir", "apps")
	wCfg.SetDefault("Description", "Add your Workspace description")
	wCfg.SetDefault("Applications", []string{})
}

// return absolute path, wd in case of - ""
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

// find first match file
func findFile(targetDir string, patterns []string) (string, error) {
	var (
		err     error
		matches []string
	)
	for _, pattern := range patterns {
		matches, err = filepath.Glob(targetDir + pattern)
		if err != nil {
			return "", fmt.Errorf("file not found")
		}

		if len(matches) != 0 {
			return matches[0], nil
		}
	}

	return "", fmt.Errorf("file not found")
}
