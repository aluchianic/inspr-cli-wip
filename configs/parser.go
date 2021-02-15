package configs

import (
	"go.uber.org/zap"
)

// parses config file according to it 'Definition'
func (w *WorkspaceFiles) Parse() *ConfigError {
	// parse workspace
	if err := w.RawConfig.parse(WorkspaceYaml{}); err != nil {
		return err
	}

	// parse application
	for _, fileRaw := range w.ApplicationsFiles {
		if err := fileRaw.parse(ApplicationYaml{}); err != nil {
			return err
		}
	}

	return nil
}

func (f *RawConfig) parse(i interface{}) *ConfigError {
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
