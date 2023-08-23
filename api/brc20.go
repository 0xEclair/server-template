package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func BRC20Balance(c *gin.Context) {
	var service service.BRC20Service
	if err := c.ShouldBind(&service); err == nil {
		res := service.AllBalance()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
