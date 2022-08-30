package main

import (
	"github.com/gin-gonic/gin"

	"server-template/api"
)

func main() {
	r := gin.Default()
	r.POST("/", api.Index)

	r.Run(":3000")
}
