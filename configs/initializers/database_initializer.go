package initializers

// DBConfiguration like postgres or mysql config
type DBConfiguration struct {
	Driver   string
	DBname   string
	Username string
	Password string
	Host     string
	Port     string
}


// InitDatabaseEnv will init configuration of DB
func InitDatabaseEnv() *DBConfiguration {
	var dbConfig DBConfiguration
	dbConfig.Driver = ""

	return &dbConfig
}
