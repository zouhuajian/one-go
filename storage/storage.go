package storage

import (
	"database/sql"
	"fmt"
)

func InitMySQL() {

	//-------1、打开数据库--------
	db, err := sql.Open("mysql", "root:112233@tcp(127.0.0.1:3305)/test?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("connect mysql failed", err)
		return
	}
	db.SetMaxIdleConns(10)
	//延迟到函数结束关闭链接
	//defer db.Close()
}
