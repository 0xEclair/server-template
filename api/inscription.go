package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

func Content(c *gin.Context) {
	var service service.InscriptionContentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find()
		//c.JSON(200, res)
		c.String(200, res.Data.(string))
	} else {
		c.JSON(200, err.Error())
	}
}

func Address(c *gin.Context) {
	var service service.ListInscriptionsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func Avatar(c *gin.Context) {
	var service service.AvatarService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func Domain(c *gin.Context) {
	var service service.DomainService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
