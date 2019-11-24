package initializers

import "gitlab.cloudwalk.work/product-center/hotpot-backend.git/configs"

// InitDatabaseEnv will init configuration of DB
func InitDatabaseEnv() *configs.DBConfiguration {
	var dbConfig configs.DBConfiguration
	dbConfig.Driver = ""

	return &dbConfig
}
