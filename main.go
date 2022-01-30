package main

import (
	"one-go/internal/config"
	"one-go/router"
)

func main() {
	conf := config.InitConfig()
	server := router.InitRouter()
	server.GinEngine.Run(":" + conf.AppConfig.Port)
}
