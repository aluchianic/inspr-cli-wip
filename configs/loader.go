package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Loads workspace and all application configs inside `WorkspaceConfig.AppsDir` and 2 level down
func (w *WorkspaceFiles) Load() *ConfigError {
	var matches []string
	// set Default
	if w.Root == "" {
		w.Root = toAbsolute("") // cwd
	}
	w.ApplicationsFiles = map[AppName]RawConfig{}

	// Load Workspace
	matches, _ = filepath.Glob(path.Join(w.Root, workspaceFileName))
	if len(matches) == 0 {
		return ErrNotFound(workspace, w.Root)
	}
	w.load(matches[0], workspace)

	// load Applications
	matches, _ = filepath.Glob(path.Join(w.Root, "**/**", applicationFileName))
	for _, match := range matches {
		name := AppName(strings.Split(path.Base(match), ".")[0])

		app := RawConfig{}
		app.load(match, application)

		w.ApplicationsFiles[name] = app
	}

	return nil
}

// Set RawConfig values
func (f *RawConfig) load(path string, definition string) {
	f.Path = path
	f.Definition = definition
	f.Config = viper.New()

	fmt.Printf("Load file: \n\t path: %s \n\t type: %s\n", f.Path, f.Definition)
}

// Returns filename for current config file
func (f *RawConfig) name() string {
	return path.Base(f.Path)
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
