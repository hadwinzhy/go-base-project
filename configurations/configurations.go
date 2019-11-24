package configurations

import (
	"fmt"

	"github.com/spf13/viper"
	"gitlab.cloudwalk.work/product-center/hotpot-backend.git/configs/initializers"
)

// Configuration will include one db config and many service config
type Configuration struct {
	Env      string
	Service  []ServiceConfiguration
	Database DBConfiguration
}

// DBConfiguration like postgres or mysql config
type DBConfiguration struct {
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
