package initializers

// ServicesConfig like kafka redis and so on...
type ServicesConfig struct {
	Name string
	Host string
	Port string
}

// InitServiceConfig will init configuration of DB
func InitServiceConfig(env string) *[]ServicesConfig {
	var serviceConfig []ServicesConfig
	// dbConfig.Driver = ""

	return &serviceConfig
}
