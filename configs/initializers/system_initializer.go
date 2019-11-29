package initializers

// SystemConfig will save all system env
type SystemConfig struct {
	Env string
}

// InitSystemConfig will init some application env
func InitSystemConfig() *SystemConfig {
	var sysConfig *SystemConfig = new(SystemConfig)
	sysConfig.Env = "dev"

	return sysConfig
}
