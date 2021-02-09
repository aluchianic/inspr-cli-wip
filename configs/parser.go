package configs

import (
	"fmt"
)

// parses config file according to it 'Definition'
func (w *WorkspaceFiles) Parse() *ConfigError {
	fileRaw := &w.FileRaw
	if err := fileRaw.parse(); err != nil {
		return err
	}

	for _, name := range w.Config.GetStringSlice("Applications") {
		fileRaw, err := w.getApp(name)
		if err != nil {
			return err
		}
		if err := fileRaw.parse(); err != nil {
			return err
		}
	}
	return nil
}

func (f *FileRaw) parse() *ConfigError {
	var i interface{}

	switch f.Definition {
	case application:
		i = ApplicationYaml{}
		break
	case workspace:
		i = WorkspaceFiles{}
		break
	}

	if err := f.unmarshal(&i); err != nil {
		return err
	}

	if err := f.read(); err != nil {
		return err
	}

	f.Parsed = true
	fmt.Printf("%s config parsed : \n\t %+v \n", f.Definition, f.Config.ConfigFileUsed())

	return nil
}

func (f *FileRaw) unmarshal(i interface{}) *ConfigError {
	if err := f.Config.Unmarshal(&i); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to unmarshal `" + f.Path + "`",
		}
	}

	return nil
}

func (f *FileRaw) read() *ConfigError {
	f.Config.SetConfigFile(f.Path)

	err := f.Config.MergeInConfig()
	if err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to merge config",
		}
	}

	if err := f.Config.ReadInConfig(); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "failed to read config",
		}
	}

	return nil
}

func (w *WorkspaceFiles) getApp(name string) (*FileRaw, *ConfigError) {
	files := w.ApplicationsFiles
	if _, ok := files[AppName(name)]; !ok {
		return nil, &ConfigError{
			Message: "Application `" + name + "` not defined in config",
		}
	}
	fileRaw := files[AppName(name)]

	return &fileRaw, nil
}
