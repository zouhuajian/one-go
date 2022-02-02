package main

import (
	"one-go/internal/config"
	"one-go/internal/cron"
	"one-go/router"
	"one-go/storage"
)

func main() {
	conf := config.InitConfig()
	server := router.InitRouter()
	storage.InitStorage(*conf)
	cron.InitCron()
	server.GinEngine.Run(":" + conf.AppConfig.Port)
}
