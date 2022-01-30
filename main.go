package main

import (
	"one-go/internal/config"
	"one-go/router"
)

func main() {
	conf := config.InitConfig()
	port := conf.AppConfig.Port
	server := router.InitRouter()
	server.GinEngine.Run(":" + port)
}
