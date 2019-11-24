package admins

import (
	"github.com/gin-gonic/gin"
	"gitlab.cloudwalk.work/product-center/hotpot-backend.git/src/admins/dao"
)

func createHandler(c *gin.Context) {
	dao.InitAdmin()
}
