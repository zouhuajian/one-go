package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"one-go/internal/config"
	"time"
)

// MySQL ...
var MySQL *gorm.DB

func InitStorage(conf config.Config) {
	initMySQL(conf.MySQLConfig)
}

func initMySQL(conf config.MySQLConfig) {
	log.Printf("mysql config {%#v}", conf)
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.UserName,
		conf.Password,
		conf.URL,
		conf.Database)
	log.Printf("connect to mysql %s", uri)
	// open mysql

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed get database")
	}

	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(10)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(10))
	MySQL = db
}
