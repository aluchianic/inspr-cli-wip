package config

import (
	"github.com/spf13/viper"
)

type Error struct {
	Err     error
	Message string
	Type    string
	Reason  string
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) AlreadyExists() bool {
	_, ok := e.Err.(viper.ConfigFileAlreadyExistsError)
	return ok
}

func (e *Error) NotFound() bool {
	_, ok := e.Err.(viper.ConfigFileNotFoundError)
	return ok
}

func ErrNotFound(definition string, path string) *Error {
	return &Error{
		Err:     viper.ConfigFileNotFoundError{},
		Message: "`" + definition + "` file not found in: " + path + "\nuse inspr init `" + definition + "`-h\n",
	}
}
