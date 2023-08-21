package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func BitmapDetail(c *gin.Context) {
	var service service.BitmapService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func BitmapCount(c *gin.Context) {
	var service service.BitmapCountService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Count()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func BitmapListByAddress(c *gin.Context) {
	address := c.Param("address")
	var service service.BitmapListService
	service.Address = address
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func BitmapRank(c *gin.Context) {
	var service service.BitmapRankService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
