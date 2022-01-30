package storage

import (
	"database/sql"
	"fmt"
	"log"
	"one-go/internal/config"
)

func InitMySQL(conf config.MySQLConfig) {
	log.Printf("mysql config {%#v}", conf)
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		conf.UserName,
		conf.Password,
		conf.URL,
		conf.Database)
	log.Printf("connect to mysql %s", uri)
	// open mysql
	db, err := sql.Open("mysql", "root:112233@tcp(127.0.0.1:3305)/test?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("connect mysql failed", err)
		return
	}
	db.SetMaxIdleConns(10)
	//延迟到函数结束关闭链接
	//defer db.Close()
}
