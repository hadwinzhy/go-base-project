package service

import (
	"github.com/gin-gonic/gin"
	"cw-app/hotpot-backend/src/admins/dao"
)

func CreateHandler(c *gin.Context) {
	dao.InitAdmin()
}
