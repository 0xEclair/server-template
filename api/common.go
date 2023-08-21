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

func Image(c *gin.Context) {
	var service service.ImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func Inscription(c *gin.Context) {
	var service service.InscriptionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
