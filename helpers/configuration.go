package helpers

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(path string) {
	locate(path)
	loadEnv()
}

func locate(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("inspr.config")
		viper.SetConfigType("yaml")

		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func loadEnv() {
	viper.AutomaticEnv()
}
