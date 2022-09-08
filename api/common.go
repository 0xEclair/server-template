package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}

func Error(c *gin.Context) {
	var service service.ErrorService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Println()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
