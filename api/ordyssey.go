package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func OrdysseyAllBitmaps(c *gin.Context) {
	var service service.AllBitmapService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}