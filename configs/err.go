package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
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

func ShowAndExistIfErrorExists(e *ConfigError) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", e.Message)
		os.Exit(1)
	}
}
