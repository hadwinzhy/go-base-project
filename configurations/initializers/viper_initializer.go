package initializers

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitViper will init viper settings
func InitViper() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
}

// ReadYAMLfile will read yaml in local
func ReadYAMLfile(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); notFound {
			fmt.Println("config not found for " + path)
		}
	}
}
