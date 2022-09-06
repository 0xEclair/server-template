package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func Ping(c *gin.Context) {
	var service service.PingService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Accept()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
