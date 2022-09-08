package server

import (
	"github.com/gin-gonic/gin"

	"server-template/api"
	"server-template/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("log", api.Error)
	}

	return r
}
