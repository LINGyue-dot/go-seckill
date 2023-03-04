package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	user_id   int
	user_name string
	user_img  string
}

func main() {
	DB, _ := sql.Open("mysql", "vze:csz51628@+@tcp(123.249.35.73:3306)/vze_db")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	rets, _ := DB.Query("select * from user")

	var users []User
	var u User
	for rets.Next() {
		err3 := rets.Scan(&u.user_id, &u.user_name, &u.user_img)
		if err3 != nil {
			fmt.Printf("Scan fail %v", err3)
			return
		}
		users = append(users, u)
	}
	for _, u := range users {
		fmt.Println(u)
	}
}
