package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"server-template/api"
	"server-template/docs"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("log", api.Error)

		v1.GET("images", api.Image)
		v1.GET("inscription", api.Inscription)
		v1.GET("content", api.Content)
		v1.GET("address", api.Address)
		v1.GET("/address/:condition", api.AddressByCondition)
		v1.GET("avatar", api.Avatar)
		v1.GET("domain", api.Domain)
	}

	bitmap := v1.Group("/collection").Group("/bitmap")
	{
		bitmap.GET("/detail", api.BitmapDetail)
		bitmap.GET("/count", api.BitmapCount)
		bitmap.GET("/:address", api.BitmapListByAddress)
		bitmap.GET("/rank", api.BitmapRank)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
