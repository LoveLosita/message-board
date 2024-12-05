package main

import (
	"fmt"
	"message-board/dao"
	"message-board/router"
)

func main() {
	err := dao.ConnectDB() //连接数据库
	if err != nil {
		fmt.Println(err)
	}
	router.RegisterRouters() //注册路由并启动服务
}
