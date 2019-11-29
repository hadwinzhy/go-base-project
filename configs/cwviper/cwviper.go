package cwviper

import (
	"os"

	"github.com/spf13/viper"
)

var configRootPath string = "configs/yaml/"

// InitViper will init viper settings
func InitViper() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
}

// ReadYAMLfile will read yaml in local
func ReadYAMLfile(path string) error {
	viper.SetConfigFile(configRootPath + path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// GetStringByEnvOrYAML will return env first or yaml sencond
func GetStringByEnvOrYAML(envName string, yamlPath string) string {
	value, isExist := os.LookupEnv(envName)
	if !isExist {
		value = viper.GetString(yamlPath)
	}
	return value
}
