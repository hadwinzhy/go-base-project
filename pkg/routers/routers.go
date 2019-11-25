package routers

import (
	"github.com/gin-gonic/gin"
	"cw-app/hotpot-backend/middleware"
	"cw-app/hotpot-backend/src/admins"
)

////////// Public Method //////////

// InitRouters :Init Routers will init all router here
func InitRouters(r *gin.Engine) {
	middleware.UseCORS(r)
	// r.Use(logger.GinLogger())
	routerErrorHandler(r)

	v1 := r.Group("/v1/api")
	routerRegister(v1)
}

////////// Private Method //////////

func routerRegister(v1 *gin.RouterGroup) {
	admins.Register(v1)
}
