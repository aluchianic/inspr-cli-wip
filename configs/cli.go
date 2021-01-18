package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

// Load CLI configurations
func LoadCli() func() {
	return func() {
		loadCache()
		loadEnv()
		loadCliConfig()

		fmt.Printf("CLI config is loaded \n Acc: %s \n Token: %s \n", viper.GetString("Acc"), viper.GetString("Token"))
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
func loadCliConfig() {
	var cName = "inspr.config"

	viper.SetConfigType("yaml")
	viper.SetConfigName(cName)

	viper.AddConfigPath(".")
	viper.AddConfigPath(cacheDir())

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			setDefaults()
			p := path.Join(cacheDir(), cName+".yaml")
			if viper.SafeWriteConfigAs(p) != nil {
				panic(fmt.Errorf("Failed to write config file: %s \n", err))
			}
		} else {
			panic(fmt.Errorf("Should not happen: %s \n", err))
		}
	}
	// CLI config file found and successfully parsed
}

// Set CLI default values
func setDefaults() {
	viper.Set("Version", "0.0.0")
	viper.Set("Acc", "123456789")
	viper.Set("Token", "aBcX-d65@-ds12")
}

// Load environment variables prefixed `INSPR_`
func loadEnv() {
	viper.SetEnvPrefix("inspr")
	viper.AutomaticEnv()
}
