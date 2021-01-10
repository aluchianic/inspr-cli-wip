package helpers

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

// Load CLI configurations
func Config() func() {
	return func() {
		loadCache()
		loadEnv()
		loadConfig()

		v := viper.GetString("Version")
		m := viper.GetString("Mode")
		c := viper.GetString("config")
		fmt.Printf("Config is loaded \n, Version: %s \n, Mode: %s \n, Config: %s \n", v, m, c)
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
func loadCache() {
	c := cacheDir()
	if err := createDirIfNotExists(c); err != nil {
		panic(fmt.Errorf("failed to create dir at %s: %s", c, err))
	}
}

// Locate and read CLI configuration file, create if not exists
func loadConfig() {
	var cfgName = "inspr.config"
	var cfgPath = viper.GetString("config")

	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName(cfgName)

		viper.AddConfigPath(".")
		viper.AddConfigPath(cacheDir())
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			setDefaults()
			p := path.Join(cacheDir(), cfgName+".yaml")
			if viper.SafeWriteConfigAs(p) != nil {
				panic(fmt.Errorf("Failed to write config file: %s \n", err))
			}
		} else {
			if cfgPath != "" {
				panic(fmt.Errorf("Config path is set incorrect: %s \n", cfgPath))
			} else {
				panic(fmt.Errorf("Should not happen: %s \n", err))
			}
		}
	}
	// Config file found and successfully parsed
}

// Set CLI default values
func setDefaults() {
	viper.SetDefault("Version", "0.0.0")
	viper.SetDefault("Mode", "production")
}

// Load environment variables prefixed `INSPR_`
func loadEnv() {
	viper.SetEnvPrefix("inspr")
	viper.AutomaticEnv()
}
