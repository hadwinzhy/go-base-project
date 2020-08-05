package configs

import "fmt"

// Configuration will include one db config and many service config
type Configuration struct {
	Services  []ServiceConfiguration
	Databases []DBConfiguration
}

// ENV is running env
var ENV string

// InitApplication will init all settings
func InitApplication() {
	initEnv()
	initConfig()
}

func initEnv() {
	readYAMLFile("configs", "application.yaml")
	ENV = readYAMLStringValue("env")
	fmt.Println(ENV)
	readYAMLFile("configs/enviorments/"+ENV, "database.yaml")
}

// initConfig will load settings
func initConfig() {
	var config Configuration
	config.Databases = initAllDBConfig()
	config.Services = initAllServiceConfig()
}

func initAllDBConfig() []DBConfiguration {
	var defaultDB DBConfiguration
	setDBConfig("db.default", &defaultDB)
	// read default

	var a []DBConfiguration
	return a
}

func initAllServiceConfig() []ServiceConfiguration {
	var a []ServiceConfiguration
	return a
}
