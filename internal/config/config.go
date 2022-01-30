package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// AppConfig ...
type AppConfig struct {
	Name string `yaml:"application"` // 自定义名称
	Port string `yaml:"port"`
}

// MySQLConfig ...
type MySQLConfig struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type config struct {
	// common
	AppConfig   AppConfig   `yaml:"application"`
	MySQLConfig MySQLConfig `yaml:"mysql"`
}

func InitConfig() *config {
	viper.SetConfigName("app")       // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	watchConfig()
	fmt.Printf("config, %v", conf)
	return conf
}

// watchConfig ...
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
