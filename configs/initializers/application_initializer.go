package initializers

import "github.com/gin-gonic/gin"


// ServiceConfiguration like kafka redis and so on...
type ServiceConfiguration struct {
	Name string
	Host string
	Port string
}

// InitApplicationEnv will init some application env
func InitApplicationEnv() {
	// ENV = os.LookupEnv("")
	// ENV := "dev"
	// if
	gin.SetMode("release")
}
