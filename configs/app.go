package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
)

var (
	aCfg = viper.New()
	app  = Application{}
)

type Channels struct {
	In  []string
	Out []string
}

// Application config
type Application struct {
	Name      AppName
	Id        string
	DependsOn []string
	Channels
}

// Creates new Application under Workspace.AppsDir/$Application.Name/
func (w *Workspace) NewApplication(name string) (*Application, *ConfigError) {
	if w.AppExists(AppName(name)) {
		return nil, &ConfigError{
			Message: "`" + name + "`" + " already exists in workspace.",
		}
	}

	var pathToApp = toAbsolute(path.Join(w.AppsDir, name))
	_ = createDirIfNotExists(w.AppsDir)
	_ = createDirIfNotExists(pathToApp)

	aCfg.Set("Name", name)
	aCfg.SetConfigName(name + ".application")
	aCfg.SetConfigType("yaml")

	aCfg.AddConfigPath(pathToApp)

	setAppDefaults()
	if err := aCfg.SafeWriteConfig(); err != nil {
		return nil, &ConfigError{
			Err:     err,
			Message: "failed to write Application config",
		}
	}

	if err := w.AddApplication(AppName(name)); err != nil {
		return nil, err
	}

	a, err := w.InitApplication(name)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Created new Application in: %s \n", aCfg.ConfigFileUsed())
	return a, err
}

// Initialize Application config by name
func (w *Workspace) InitApplication(name string) (*Application, *ConfigError) {
	if !w.AppExists(AppName(name)) {
		return nil, &ConfigError{
			Message: "Passed name doesn't exist in config",
		}
	}

	n := name + ".application.yaml"
	patterns := []string{n, "**/" + n, "**/**/" + n}
	confPath, err := findFile(w.Root(), patterns)
	if err != nil {
		fmt.Printf("ERRROR : %s \n", err)
	}

	aCfg.SetConfigFile(confPath)
	if err := aCfg.ReadInConfig(); err != nil {
		return nil, &ConfigError{
			Err:     err,
			Message: "failed to read config",
		}
	}

	if err := aCfg.Unmarshal(&workspaceConf); err != nil {
		return nil, &ConfigError{
			Message: "unable to decode into struct",
			Err:     err,
		}
	}

	return &app, nil
}

// Prints all currently initialized Application settings
func (a *Application) Describe() *ConfigError {
	if aCfg.ConfigFileUsed() == "" {
		return &ConfigError{
			Err:     viper.ConfigFileNotFoundError{},
			Message: "can't describe, application config is not located. Use inspr init [name] -a to create new Application",
		}
	}
	fmt.Printf("Application config used: %s \n", aCfg.ConfigFileUsed())
	return nil
}

// Sets Application default values
func setAppDefaults() {
	aCfg.SetDefault("Depends", []string{})
	aCfg.SetDefault("Channels", &Channels{})
	aCfg.SetDefault("Description", "Add your Application description")
}
