package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func GetBitmap(c *gin.Context) {
	var service service.InscriptionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
