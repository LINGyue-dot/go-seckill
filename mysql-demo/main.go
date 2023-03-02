package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB, _ := sql.Open("mysql", "${username}:${password}@tcp(123.249.35.73:3306)/vze_db")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	rows, e := DB.Query("select * from user")
	fmt.Println("connnect success")
}
