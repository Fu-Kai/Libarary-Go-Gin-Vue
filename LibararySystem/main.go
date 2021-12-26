package main

import (
	"LibararySystem/db"
	"LibararySystem/handler"
	"fmt"
)

func main() {
	// 初始化mysql
	err := db.Init()
	if err != nil {
		fmt.Printf("初始化数据库错误:%v\n", err)
		return
	}

	// 初始化web服务器，监听的地址格式为 0.0.0.0:8080
	err = handler.Start()
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)
		return
	}
}
