package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"message-board/model"
	"message-board/service"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	postUser := model.User{}
	err := c.BindJSON(&postUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	result := false
	result, err = service.UserLogin(postUser.UserName, postUser.PassWord) //调用用户登录模块
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if result {
		c.JSON(consts.StatusOK, map[string]string{
			"status": "success",
		})
	} else {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "wrong-pwd",
		})
	}
}

func UserRegister(ctx context.Context, c *app.RequestContext) {
	user := model.User{}
	err := c.BindJSON(&user) //解析获取的JSON，存入结构体
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	err = service.UserRegister(user) //调用用户注册函数，传入用户数据完成注册
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, map[string]string{
		"状态": "成功插入",
	})
}

//以下是管理员专属功能

func ShowUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}

func DeleteUser(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}
