package helpers

import (
	"fmt"
	"github.com/spf13/viper"
)

// Load CLI configurations
func Config(path string) func() {
	return func() {
		loadConfig(path)
		loadEnv()

		v := viper.GetString("Version")
		m := viper.GetString("Mode")
		fmt.Printf("Config is loaded \n, Version: %s \n, Mode: %s \n", v, m)
	}
}

// Locate and read CLI configuration file, create if not exists
func loadConfig(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("inspr.config")
		viper.SetConfigType("yaml")

		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.inspr")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			setDefaults()
			if viper.SafeWriteConfig() != nil {
				panic(fmt.Errorf("Failed to write config file: %s \n", err))
			}
		} else {
			panic(fmt.Errorf("Config file was found but another error was produced: %s \n", err))
		}
	}
	// Config file found and successfully parsed
}

// Set CLI default values
func setDefaults() {
	viper.SetDefault("Version", "0.0.0")
	viper.SetDefault("Mode", "production")
}

// Load environment variables prefixed `INSPR`
func loadEnv() {
	viper.SetEnvPrefix("inspr")
	viper.AutomaticEnv()
}
