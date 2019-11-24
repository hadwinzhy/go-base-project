package service

import (
	"github.com/gin-gonic/gin"
	"gitlab.cloudwalk.work/product-center/pc-public/cw-app/hotpot-backend.git/src/admins/dao"
)

func CreateHandler(c *gin.Context) {
	dao.InitAdmin()
}
