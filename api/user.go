package api

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"message-board/auth"
	"message-board/model"
	"message-board/service"
	"message-board/utils"
	"strconv"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	postUser := model.User{}
	err := c.BindJSON(&postUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	result := false
	userID := 0
	userID, result, err = service.UserLogin(postUser.UserName, postUser.PassWord) //调用用户登录模块
	if err != nil {
		switch {
		case errors.Is(err, utils.WrongUsrName): //用户名错误
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	if result {
		// 创建 JWT
		strJWT, err := auth.GenerateJWT(userID)
		if err != nil {
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		}
		combinedJson := map[string]interface{}{ //返回token
			"token":    strJWT,
			"response": utils.Ok,
		}
		c.JSON(consts.StatusOK, combinedJson)
	} else {
		c.JSON(consts.StatusBadRequest, utils.WrongPwd) //密码错误
	}
}

func UserRegister(ctx context.Context, c *app.RequestContext) {
	user := model.User{}
	err := c.BindJSON(&user) //解析获取的JSON，存入结构体
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	err = service.UserRegister(user) //调用用户注册函数，传入用户数据完成注册
	if err != nil {
		if errors.Is(err, utils.MissingParam) || errors.Is(err, utils.InvalidUsername) { //客户端原因的错误
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		} else {
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	c.JSON(consts.StatusOK, utils.Ok)
}

//以下是管理员专属功能

func ShowUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能
	inquireJson := model.JsonInquiry{} //获取前端传来的json
	err := c.BindJSON(&inquireJson)    //解析json
	if err != nil {                    //解析失败
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	handlerID := int(c.GetFloat64("user_id"))                                              //获取用户ID
	user, err := service.ShowUserInfo(inquireJson.UserName, inquireJson.UserID, handlerID) //调用service层的方法，获取用户信息
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUnauthorized):
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		case errors.Is(err, utils.InvalidID):
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		case errors.Is(err, utils.WrongUsrName):
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	combinedJson := map[string]interface{}{
		"user":         user,
		"respond code": utils.Ok,
	}
	c.JSON(consts.StatusOK, combinedJson) //返回查询结果
}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能
	handlerID := int(c.GetFloat64("user_id")) //获取用户ID
	jsonUser := model.NewUser{}
	err := c.BindJSON(&jsonUser) //解析json
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	err = service.ChangeUserInfo(handlerID, jsonUser) //调用service层的方法，修改用户信息
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUnauthorized):
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	c.JSON(consts.StatusOK, utils.Ok) //返回查询结果
}

func DeleteUser(ctx context.Context, c *app.RequestContext) { //后期再添加的功能
	handlerID := int(c.GetFloat64("user_id")) //获取用户ID
	targetID := c.Query("id")
	if targetID == "" { //缺少参数
		c.JSON(consts.StatusBadRequest, utils.MissingParam)
		return
	}
	intTargetID, err := strconv.ParseInt(targetID, 10, 0) //解析参数
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err)) //解析失败
		return
	}
	err = service.DeleteUser(int(intTargetID), handlerID) //调用service层的方法，删除用户
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUnauthorized): //不是管理员，无法删除
			c.JSON(consts.StatusBadRequest, utils.ClientError(err)) //返回错误
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err)) //内部错误
		}
		return
	}
	c.JSON(consts.StatusOK, utils.Ok) //返回成功
}
