package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

type CliConfig struct {
	Version string
	Account string
	Token   string
}

func InitCLi() {
	loadCliCache()
	loadCliEnv()
	if err := loadCliConfig(); err != nil {
		os.Exit(1)
	}
}

// Creates directory
func createDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(path, 0754); err != nil {
			return err
		}
	}
	return nil
}

// Returns path to $HOME/.inspr folder
func cacheDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get cache dir: %s", err))
	}

	return path.Join(home, ".inspr")
}

// Load cache ...
func loadCliCache() {
	c := cacheDir()
	if err := createDirIfNotExists(c); err != nil {
		panic(fmt.Errorf("failed to create dir at %s: %s", c, err))
	}
}

// Locate and read CLI configuration file, create if not exists
func loadCliConfig() *ConfigError {
	var cliConf CliConfig

	viper.SetConfigType("yaml")
	viper.SetConfigName("inspr.config")

	viper.AddConfigPath(cacheDir())
	setCliDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if viper.SafeWriteConfig() != nil {
				return &ConfigError{
					Err:     err,
					Message: "failed to write CLI config file",
				}
			}
		}
	}

	if err := viper.Unmarshal(&cliConf); err != nil {
		return &ConfigError{
			Err:     err,
			Message: "unable to decode CLI config into struct",
		}
	}

	fmt.Printf("cliConf :: %+v", cliConf)

	return nil
}

// Set CLI default values
func setCliDefaults() {
	viper.Set("Version", "0.0.0")
	viper.Set("Account", "123456789")
	viper.Set("Token", "aBcX-d65@-ds12")
}

// Load environment variables prefixed `INSPR_`
func loadCliEnv() {
	viper.SetEnvPrefix("inspr")
	viper.AutomaticEnv()
}
