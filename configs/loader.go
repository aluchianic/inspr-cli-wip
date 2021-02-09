package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Set FileRaw values
func (f *FileRaw) load(path string, definition string) {
	f.Path = path
	f.Definition = definition
	f.Config = viper.New()

	fmt.Printf("Load file: \n\t path: %s \n\t type: %s\n", f.Path, f.Definition)

}

// Loads workspace and all application configs inside `WorkspaceConfig.AppsDir` and 2 level down
func (w *WorkspaceFiles) Load() *ConfigError {
	var matches []string
	// set Default
	if w.Root == "" {
		w.Root = toAbsolute("") // cwd
	}
	w.ApplicationsFiles = map[AppName]FileRaw{}

	// Load Workspace
	pattern := "/" + workspaceFileName
	matches, _ = filepath.Glob(path.Join(w.Root + pattern))
	if len(matches) == 0 {
		return &ConfigError{
			Err: errors.New("file doesn't exist"),
		}
	}
	w.load(matches[0], workspace)

	// load Applications
	appPattern := "/" + applicationFileName
	matches, _ = filepath.Glob(path.Join(w.Root, "**/**", appPattern))
	for _, match := range matches {
		name := AppName(strings.Split(path.Base(match), ".")[0])

		app := FileRaw{}
		app.load(match, application)

		w.ApplicationsFiles[name] = app
	}

	return nil
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
