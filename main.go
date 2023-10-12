package main

import (
	"server-template/config"
	"server-template/server"
)

func main() {
	config.Init()

	r := server.NewRouter()
	r.Run(":3456")
}
