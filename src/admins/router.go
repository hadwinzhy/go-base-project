package admins

import (
	"github.com/gin-gonic/gin"
	"gitlab.cloudwalk.work/product-center/pc-public/cw-app/hotpot-backend.git/src/admins/service"
)

////////// Public Methods //////////

// Register will register routers
func Register(r *gin.RouterGroup) {
	r.GET("/admin", service.CreateHandler)
}

////////// Private Methods //////////
