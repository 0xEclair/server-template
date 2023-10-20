package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func AssetsListByAddress(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithOss()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AudioByAddress(c *gin.Context) {
	var service service.AudioService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func ModelByAddress(c *gin.Context) {
	var service service.AudioService
	service.ContentType = "model"
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListModel()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func CanMergeByAddress(c *gin.Context) {
	var service service.AudioService
	service.ContentType = "model"
	if err := c.ShouldBind(&service); err == nil {
		res := service.CanMerge()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
