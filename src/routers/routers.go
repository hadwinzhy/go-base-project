package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.cloudwalk.work/product-center/hotpot-backend.git/middleware"
)

// InitRouters :Init Routers will init all router here
func InitRouters(r *gin.Engine) {
	middleware.UseCORS(r)
}
