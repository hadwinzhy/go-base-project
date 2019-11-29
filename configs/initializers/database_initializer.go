package initializers

import (
	"cw-app/hotpot-backend/configs/cwviper"
	"strings"

	"github.com/spf13/viper"
)

// DatabaseConfig like postgres or mysql config
type DatabaseConfig struct {
	Driver   string
	DBname   string
	Username string
	Password string
	Host     string
	Port     string
	SslMode  string
}

// InitDatabaseConfig will init configuration of DB
func InitDatabaseConfig(env string) *DatabaseConfig {
	driver := strings.ToUpper(viper.GetString(env + ".db.driver"))
	dbConfig := DatabaseConfig{
		Driver:   viper.GetString(env + ".db.driver"),
		DBname:   viper.GetString(env + ".db.dbname"),
		Username: cwviper.GetStringByEnvOrYAML(driver+"_DBUSER", env+".db.user"),
		Password: cwviper.GetStringByEnvOrYAML(driver+"_DBPASSWORD", env+".db.password"),
		Host:     cwviper.GetStringByEnvOrYAML(driver+"_DBHOST", env+".db.host"),
		Port:     cwviper.GetStringByEnvOrYAML(driver+"_DBPORT", env+".db.port"),
		SslMode:  viper.GetString(env + ".db.sslmode"),
	}

	return &dbConfig
}
