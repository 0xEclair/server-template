package api

import (
	"github.com/gin-gonic/gin"

	"server-template/service"
)

// @BasePath /api/v1

// BitmapDetail godoc
// @Summary bitmap信息
// @Schemes
// @Description bitmap detail
// @Tags bitmap
// @Accept application/json
// @Produce application/json
// @Param id query []int true "id"
// @Success 200 {object} serializer.Response{data=[]model.Bitmap}
// @Router /collection/bitmap/detail [get]
func BitmapDetail(c *gin.Context) {
	var service service.BitmapService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

// @BasePath /api/v1

// BitmapCount godoc
// @Summary bitmap总数
// @Schemes
// @Description bitmap count
// @Tags bitmap
// @Accept application/json
// @Produce application/json
// @Success 200 {object} serializer.Response{data=int32}
// @Router /collection/bitmap/count [get]
func BitmapCount(c *gin.Context) {
	var service service.BitmapCountService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Count()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

// @BasePath /api/v1

// BitmapListByAddress godoc
// @Summary 地址下所有bitmap
// @Schemes
// @Description bitmap list of address
// @Tags bitmap
// @Accept application/json
// @Produce application/json
// @Param address query string true "address"
// @Param offset query int32 false "offset" default(0)
// @Param limit query int32 false "limit" default(20)
// @Success 200 {object} serializer.Response{data=[]model.Bitmap}
// @Router /collection/bitmap/:address [get]
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

// @BasePath /api/v1

// BitmapRankExample godoc
// @Summary holders排名
// @Schemes
// @Description rank for bitmap holders
// @Tags bitmap
// @Accept application/json
// @Produce application/json
// @Param offset query int32 false "offset" default(0)
// @Param limit query int32 false "limit" default(20)
// @Success 200 {object} serializer.Response{data=[]model.BitmapRank}
// @Router /collection/bitmap/rank [get]
func BitmapRank(c *gin.Context) {
	var service service.BitmapRankService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
