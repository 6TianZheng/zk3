package service

import (
	"net/http"
	"zk3/bff/basic/config"
	__ "zk3/bff/basic/proto"
	"zk3/bff/handler/request"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var form request.Login
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数有误",
			"msg":   400,
		})
		return
	}

	l, err := config.GoodsClient.Login(c, &__.LoginReq{
		Username: form.Username,
		Password: form.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数有误",
			"msg":   400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": l.Msg,
		"msg":   l.Code,
	})
	return
}
