package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// UseCORS will set cors config
func UseCORS(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	config.AddAllowHeaders("X-Requested-With")

	config.AddAllowMethods("OPTIONS")
	config.AddAllowMethods("DELETE")
	config.AddAllowMethods("PATCH")

	r.Use(cors.New(config))
}
