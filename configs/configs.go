package configs

import (
	"cw-app/hotpot-backend/configs/cwviper"
	"cw-app/hotpot-backend/configs/initializers"
)

// SystemConfig will return system config
var SystemConfig *initializers.SystemConfig

// ServicesConfig will save database config
var ServicesConfig *[]initializers.ServicesConfig

// DatabaseConfig will save database config
var DatabaseConfig *initializers.DatabaseConfig

// InitConfig will load settings
func InitConfig() {

	// Step 1 Read all yaml in system
	cwviper.InitViper()
	if err := cwviper.ReadYAMLfile("system.yaml"); err != nil {
		panic("config yaml failed: " + err.Error())
	}
	SystemConfig = initializers.InitSystemConfig()

	// Step 2 Read all yaml in service
	if err := cwviper.ReadYAMLfile("service.yaml"); err != nil {
		panic("config yaml failed: " + err.Error())
	}
	ServicesConfig = initializers.InitServiceConfig(SystemConfig.Env)

	// Step 3 Read all yaml in database
	if err := cwviper.ReadYAMLfile("database.yaml"); err != nil {
		panic("config yaml failed: " + err.Error())
	}
	DatabaseConfig = initializers.InitDatabaseConfig(SystemConfig.Env)
}
