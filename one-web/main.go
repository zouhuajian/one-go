package main

import (
	"github.com/one-go/one-web/internal/config"
	"github.com/one-go/one-web/internal/cron"
	"github.com/one-go/one-web/router"
	"github.com/one-go/one-web/storage"
)

func main() {
	conf := config.InitConfig()
	server := router.InitRouter()
	storage.InitStorage(*conf)
	cron.InitCron()
	server.GinEngine.Run(":" + conf.AppConfig.Port)
}
