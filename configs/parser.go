package configs

import "go.uber.org/zap"

// parses config file according to it 'Definition'
func (w *WorkspaceFiles) Parse() *ConfigError {
	fileRaw := w.RawConfig
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

func (w *WorkspaceFiles) getApp(name string) (*RawConfig, *ConfigError) {
	if name == "" {
		return &w.RawConfig, nil
	}

	files := w.ApplicationsFiles
	if _, ok := files[AppName(name)]; !ok {
		return nil, &ConfigError{
			Message: "Application `" + name + "` not defined in config",
		}
	}
	fileRaw := files[AppName(name)]

	return &fileRaw, nil
}

func (f *RawConfig) parse() *ConfigError {
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

	f.Logger.Info("Parsed config", zap.String("path", f.Config.ConfigFileUsed()), zap.String("type", f.Definition))

	return nil
}

func (f *RawConfig) unmarshal(i interface{}) *ConfigError {
	if err := f.Config.Unmarshal(&i); err != nil {
		f.Logger.Error("failed to unmarshal", zap.String("path", f.Config.ConfigFileUsed()), zap.String("type", f.Definition))
	}

	return nil
}

func (f *RawConfig) read() *ConfigError {
	f.Config.SetConfigFile(f.Path)

	if err := f.Config.MergeInConfig(); err != nil {
		f.Logger.Error("failed to merge config", zap.String("path", f.Config.ConfigFileUsed()), zap.String("type", f.Definition))
	}
	if err := f.Config.ReadInConfig(); err != nil {
		f.Logger.Error("failed to read config", zap.String("path", f.Config.ConfigFileUsed()), zap.String("type", f.Definition))
	}

	return nil
}
