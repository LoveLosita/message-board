package cmd

import (
	"fmt"
	"message-board/api"
	"message-board/dao"
)

func main() {
	api.RegisterRouters()  //注册路由
	err := dao.ConnectDB() //连接数据库
	if err != nil {
		fmt.Println(err)
	}
	defer dao.DisconnectDB() //在退出时断开数据库连接
}
