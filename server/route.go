package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"server-template/api"
	"server-template/docs"
	"server-template/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("log", api.Error)

		v1.GET("images", api.Image)
		v1.GET("imagesv2", api.Imagev2)
		v1.GET("audio", api.AudioByAddress)
		v1.GET("model", api.ModelByAddress)
		v1.GET("merge", api.CanMergeByAddress)
		v1.GET("inscription", api.Inscription)
		v1.GET("content", api.Content)
		v1.GET("/:address", api.Address)
		v1.GET("/address/:condition", api.AddressByCondition)
		v1.GET("avatar", api.Avatar)
		v1.GET("avatarv2", api.Avatarv2)
		v1.GET("domain", api.Domain)
		v1.GET("asset", api.AssetsListByAddress)
		v1.GET("assetv2", api.AssetsListByAddressV2)
		v1.GET("assetv3", api.AssetsListByAddressV3)
		v1.GET("assetwithdlc", api.AssetsListByAddressV4)
		v1.GET("assetcollection", api.AssetsCollection)
		v1.GET("metaversedlc", api.AssetsDLCListByAddress)
		v1.GET("/assumerole", api.CreateOssKey)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("domain", api.Domainv2)
	}

	bitmap := v1.Group("/collection").Group("/bitmap")
	{
		bitmap.GET("/detail", api.BitmapDetail)
		bitmap.GET("/count", api.BitmapCount)
		bitmap.GET("/:address", api.BitmapListByAddress)
		bitmap.GET("/rank", api.BitmapRank)
		bitmap.GET("/list", api.OrdysseyAllBitmaps)
	}

	brc20 := v1.Group("/brc20")
	{
		brc20.GET("/balance", api.BRC20Balance)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
