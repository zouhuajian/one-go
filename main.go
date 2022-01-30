package main

import (
	"one-go/internal/config"
	"one-go/router"
	"one-go/storage"
)

func main() {
	conf := config.InitConfig()
	server := router.InitRouter()
	storage.InitStorage(*conf)
	server.GinEngine.Run(":" + conf.AppConfig.Port)
}
