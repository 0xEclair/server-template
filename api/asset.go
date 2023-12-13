package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func AssetsListByAddress(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithOssAndBRC420()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsListByAddressV2(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithOssAndBRC420V2()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsListByAddressV3(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithOssAndBRC420V3()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsListByAddressV42(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithOssAndBRC420V5()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsListByAddressV4(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithOssAndBRC420V4()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsCollection(c *gin.Context) {
	var service service.AssetsCollectionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Collection()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsCollectionV2(c *gin.Context) {
	var service service.AssetsCollectionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CollectionV2()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func CollectionDLC(c *gin.Context) {
	var service service.CollectionDLCService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func AssetsDLCListByAddress(c *gin.Context) {
	var service service.AssetsListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListDLCWithOssAndBRC420()
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
