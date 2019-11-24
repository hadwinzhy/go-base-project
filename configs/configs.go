package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"gitlab.cloudwalk.work/product-center/pc-public/cw-app/hotpot-backend.git/configs/initializers"
)

// Configuration will include one db config and many service config
type Configuration struct {
	Env      string
	Service  []initializers.ServiceConfiguration
	Database initializers.DBConfiguration
}

// AppConfig is system singleton config
// var AppConfig *Configuration

// InitConfig will load settings
func InitConfig() {
	initializers.InitViper()

	// set application yaml & vairable
	initializers.ReadYAMLfile("application.yaml")
	initializers.InitApplicationEnv()

	// set database yaml & vairable
	initializers.ReadYAMLfile("database.yaml")
	initializers.InitDatabaseEnv()

	host := viper.GetString("dev.db.host")
	fmt.Println("host is " + host)
}
