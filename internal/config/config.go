package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// AppConfig ...
type AppConfig struct {
	Name string `mapstructure:"application"` // 自定义名称
	Port string `mapstructure:"port"`
}

// MySQLConfig ...
type MySQLConfig struct {
	URL      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database int    `mapstructure:"database"`
}

type Config struct {
	// common
	AppConfig   AppConfig   `mapstructure:"application"`
	MySQLConfig MySQLConfig `mapstructure:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis"`
}

func InitConfig() *Config {
	viper.SetConfigName("app")       // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into config struct, %w \n", err))
	}
	log.Printf("config: %v", conf)
	watchConfig()
	return conf
}

// watchConfig ...
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
