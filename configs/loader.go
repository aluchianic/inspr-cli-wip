package configs

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func (w *WorkspaceFiles) Load(entity string) *ConfigError {
	var matches []string

	switch entity {

	case "workspace":
		pattern := "/" + workspaceFileName
		root := toAbsolute("")
		w.Definition = "workspace"

		matches, _ = filepath.Glob(path.Join(root + pattern))

		if len(matches) == 0 {
			w.Path = root
			return &ConfigError{
				Err: errors.New("file doesn't exist"),
			}
		} else {
			w.Path = matches[0]
		}
		return nil
	case "applications":
		pattern := "/" + applicationFileName

		if w.Path == "" {
			return &ConfigError{
				Err: errors.New("can't load application outside Workspace"),
			}
		}
		matches, _ = filepath.Glob(path.Join(w.Path, "**/**", pattern))

		for _, p := range matches {
			app := FileRaw{
				Path:       p,
				Definition: "application",
			}

			w.Apps = append(w.Apps, app)
		}
		return nil
	}
	return &ConfigError{
		Err: errors.New("failed to load " + entity + "  config['s]"),
	}
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
