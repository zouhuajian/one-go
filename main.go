package main

import (
	"github.com/spf13/viper"
	"one-go/internal/config"
	"one-go/router"
)

func main() {
	config.InitConfig()
	port := viper.GetString("application.port")
	server := router.InitRouter()
	server.GinEngine.Run(":" + port)
}
