package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// MySQLConfig ...
type MySQLConfig struct {
	Name     string `json:"name"` // 自定义名称
	URL      string `json:"url"`
	Database string `json:"database"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func InitConfig() {
	viper.SetConfigName("app")       // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
