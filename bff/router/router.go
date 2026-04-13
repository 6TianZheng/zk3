package router

import (
	"zk3/bff/handler/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/login", service.Login)

	return r
}
