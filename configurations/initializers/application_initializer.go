package initializers

import "github.com/gin-gonic/gin"

// InitApplicationEnv will init some application env
func InitApplicationEnv() {
	// ENV = os.LookupEnv("")
	// ENV := "dev"
	// if
	gin.SetMode("release")
}
