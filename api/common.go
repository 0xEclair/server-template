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

// @BasePath /api/v1

// Image godoc
// @Summary 查询所有图片
// @Schemes
// @Description query all images
// @Tags other
// @Accept application/json
// @Produce application/json
// @Param offset query int32 false "offset" default(0)
// @Param limit query int32 false "limit" default(0)
// @Success 200 {object} serializer.Response{data=[]model.Inscription}
// @Router /images [get]
func Image(c *gin.Context) {
	var service service.ImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func ImageId(c *gin.Context) {
	var service service.ImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListV2()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func Imagev2(c *gin.Context) {
	var service service.ImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListWithHTML()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

// @BasePath /api/v1

// InscriptionDetail godoc
// @Summary 铭文信息
// @Schemes
// @Description inscription detail
// @Tags other
// @Accept application/json
// @Produce application/json
// @Param number query int32 false "number" default(0)
// @Param id query string false "id"
// @Success 200 {object} serializer.Response{data=model.Inscription}
// @Router /inscription [get]
func Inscription(c *gin.Context) {
	var service service.InscriptionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

// @BasePath /api/v1

// AddressByCondition godoc
// @Summary 域名, number, 铭文id查铭文信息
// @Schemes
// @Description inscription detail
// @Tags other
// @Accept application/json
// @Produce application/json
// @Param p path string true "任意参数"
// @Success 200 {object} serializer.Response{data=model.Inscription}
// @Router /address/{p} [get]
func AddressByCondition(c *gin.Context) {
	var service service.AddressByConditionService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Find()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func CreateOssKey(c *gin.Context) {
	var service service.OssService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.CreateKey()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
