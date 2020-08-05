package configs

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func readYAMLFile(path string, filePath string) {
	// name := strings.TrimSuffix(fn, path.Ext(fn))

	// viper.SetConfigName(filepath.)
	viper.SetConfigType(filepath.Ext(filePath))
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
}

func readYAMLStringValue(key string) string {
	return viper.GetString(key)
}
