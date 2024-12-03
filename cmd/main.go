package main

import (
	"fmt"
	"message-board/api"
	"message-board/dao"
)

func main() {
	err := dao.ConnectDB() //连接数据库
	if err != nil {
		fmt.Println(err)
	}
	api.RegisterRouters()    //注册路由并启动服务
	defer dao.DisconnectDB() //在退出时断开数据库连接
}
