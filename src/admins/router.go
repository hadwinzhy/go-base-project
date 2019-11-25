package admins

import (
	"github.com/gin-gonic/gin"
	"cw-app/hotpot-backend/src/admins/service"
)

////////// Public Methods //////////

// Register will register routers
func Register(r *gin.RouterGroup) {
	r.GET("/admin", service.CreateHandler)
}

////////// Private Methods //////////
