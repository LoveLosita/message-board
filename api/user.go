package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"message-board/model"
	"message-board/service"
	"message-board/utils"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	postUser := model.User{}
	err := c.BindJSON(&postUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.CustomError(err))
		return
	}
	result := false
	result, err = service.UserLogin(postUser.UserName, postUser.PassWord) //调用用户登录模块
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.CustomError(err))
		return
	}
	if result {
		c.JSON(consts.StatusOK, utils.Ok)
	} else {
		c.JSON(consts.StatusBadRequest, utils.WrongPwd)
	}
}

func UserRegister(ctx context.Context, c *app.RequestContext) {
	user := model.User{}
	err := c.BindJSON(&user) //解析获取的JSON，存入结构体
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.CustomError(err))
		return
	}
	err = service.UserRegister(user) //调用用户注册函数，传入用户数据完成注册
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.CustomError(err))
		return
	}
	c.JSON(consts.StatusOK, utils.CustomSuccess("成功插入数据"))
}

//以下是管理员专属功能

func ShowUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}

func DeleteUser(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}
