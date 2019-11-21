package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration will include one db config and many service config
type Configuration struct {
	Service  []ServiceConfiguration
	Database DatabaseConfiguration
}

// DatabaseConfiguration like postgres or mysql config
type DatabaseConfiguration struct {
	Driver   string
	DBname   string
	Username string
	Password string
	Host     string
	Port     string
}

// ServiceConfiguration like kafka redis and so on...
type ServiceConfiguration struct {
	Name string
	Host string
	Port string
}

// InitConfig will load settings
func InitConfig() {
	// var configuration Configuration
}

func readYAMLfile(path string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); notFound {
			fmt.Println("config not found for " + path)
		} else {
			fmt.Println()
		}
	}
}
