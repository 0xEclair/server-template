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

// @BasePath /api/v1

// 地址下所有铭文 godoc
// @Summary 地址下所有铭文
// @Schemes
// @Description inscriptions of address
// @Tags other
// @Accept application/json
// @Produce application/json
// @Param address path string true "address"
// @Param offset query int32 false "offset" default(0)
// @Param limit query int32 false "limit" default(20)
// @Success 200 {object} serializer.Response{data=serializer.ImageListResponse}
// @Router /{address} [get]
func Address(c *gin.Context) {
	var service service.ListInscriptionsService
	if err := c.ShouldBindUri(&service); err == nil {

	} else {
		c.JSON(200, err.Error())
	}

	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

// @BasePath /api/v1

// Avatar godoc
// @Summary 地址所有图片信息
// @Schemes
// @Description query all avatars
// @Tags other
// @Accept application/json
// @Produce application/json
// @Param address query string true "address"
// @Param offset query int32 false "offset" default(0)
// @Param limit query int32 false "limit" default(0)
// @Success 200 {object} serializer.Response{data=[]model.Inscription}
// @Router /avatar [get]
func Avatar(c *gin.Context) {
	var service service.AvatarService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func Avatarv2(c *gin.Context) {
	var service service.AvatarService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithHTML()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

// @BasePath /api/v1

// Domain godoc
// @Summary 地址所有域名信息
// @Schemes
// @Description query all domains
// @Tags other
// @Accept application/json
// @Produce application/json
// @Param address query string true "address"
// @Success 200 {object} serializer.Response{data=[]model.Inscription}
// @Router /domain [get]
func Domain(c *gin.Context) {
	var service service.DomainService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
