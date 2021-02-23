package config

import (
	"github.com/spf13/viper"
)

type ConfigError struct {
	Err     error
	Message string
}

func (e *ConfigError) Error() string {
	return e.Err.Error()
}

func (e *ConfigError) AlreadyExists() bool {
	_, ok := e.Err.(viper.ConfigFileAlreadyExistsError)
	return ok
}

func (e *ConfigError) NotFound() bool {
	_, ok := e.Err.(viper.ConfigFileNotFoundError)
	return ok
}

func ErrNotFound(definition string, path string) *ConfigError {
	return &ConfigError{
		Err:     viper.ConfigFileNotFoundError{},
		Message: "`" + definition + "` file not found in: " + path + "\nuse inspr init `" + definition + "`-h\n",
	}
}
