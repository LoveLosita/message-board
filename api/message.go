package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"message-board/model"
	"message-board/service"
	"message-board/utils"
	"strconv"
)

func SendMessage(ctx context.Context, c *app.RequestContext) {
	message := model.Message{}
	err := c.BindJSON(&message)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	if message.Content == "" {
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.MissingParam))
		return
	}
	getID := c.GetFloat64("user_id") //从上下文中获取用户的id
	if getID == 0 {                  //id空白，表示未登录，不过一般在中间件就会被截止了
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.NotLoggedIn))
	}
	id := int(getID)
	fmt.Println(id)
	message.UserID = id
	err = service.SendMessage(message)
	if err != nil {
		switch {
		case errors.Is(err, utils.InvalidID):
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	c.JSON(consts.StatusOK, utils.Ok)
}

// 下面都是管理员专属功能

func GetAllMessages(ctx context.Context, c *app.RequestContext) { //获取所有评论
	handlerID := int(c.GetFloat64("user_id"))             //从上下文中获取用户的id
	allComments, err := service.GetAllMessages(handlerID) //调用service层的方法，获取所有评论
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.ServerError(err)) //如果出错，说明是服务器错误，返回500
		return
	}
	combinedJson := map[string]interface{}{ //将所有评论和状态码组合成一个json
		"messages":     allComments,
		"respond code": utils.Ok,
	}
	c.JSON(consts.StatusOK, combinedJson) //返回所有评论和状态码
}

func DeleteMessage(ctx context.Context, c *app.RequestContext) { //管理员专属功能,传入ID以删除评论
	messageID := c.Query("id")
	handlerID := int(c.GetFloat64("user_id"))
	if messageID == "" { //如果没有传入ID，那么返回错误，避免下方的转换出错
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.MissingParam))
		return
	}
	intMsgID, err := strconv.ParseInt(messageID, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	err = service.DeleteMessages(int(intMsgID), handlerID) //调用service层的方法，并传入ID进行管理员身份验证，然后删除
	if err != nil {
		switch {
		case errors.Is(err, utils.CantFindMessage):
			c.JSON(consts.StatusBadRequest, utils.NotFoundError(err))
		case errors.Is(err, utils.ErrUnauthorized):
			c.JSON(consts.StatusUnauthorized, utils.ClientError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	c.JSON(consts.StatusOK, utils.Ok)
}

func SearchForMessages(ctx context.Context, c *app.RequestContext) { //管理员专属功能
	searchParams := model.SearchParams{}
	handlerID := c.GetFloat64("user_id") //从上下文中获取用户的id
	err := c.BindJSON(&searchParams)     //绑定前端传来的参数
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err)) //参数错误
		return
	}
	comments, err := service.SearchForMessages(searchParams.CommentID, searchParams.Content, searchParams.UserID,
		searchParams.Username, int(handlerID)) //调用service层的方法，进行查询
	if err != nil {
		switch {
		case errors.Is(err, utils.MissingParam), errors.Is(err, utils.ErrUnauthorized), errors.Is(err, utils.InvalidID), errors.Is(err, utils.CantFindMessage): //参数不足
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		default: //其他错误
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	combinedJson := map[string]interface{}{
		"messages":     comments,
		"respond code": utils.Ok,
	}
	c.JSON(consts.StatusOK, combinedJson) //返回查询结果
}

func LikeMessage(ctx context.Context, c *app.RequestContext) {
	messageID := c.Query("id")
	if messageID == "" {
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.MissingParam))
		return
	}
	intMessageID, err := strconv.ParseInt(messageID, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	err = service.LikeMessage(int(intMessageID), handlerID)
	if err != nil {
		switch {
		case errors.Is(err, utils.CantFindMessage), errors.Is(err, utils.MessageAlreadyLiked):
			c.JSON(consts.StatusBadRequest, utils.ClientError(utils.CantFindMessage))
			return
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, utils.Ok)
}

func DislikeMessage(ctx context.Context, c *app.RequestContext) {
	messageID := c.Query("id")
	if messageID == "" {
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.MissingParam))
		return
	}
	intMessageID, err := strconv.ParseInt(messageID, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	err = service.DislikeMessage(int(intMessageID), handlerID)
	if err != nil {
		switch {
		case errors.Is(err, utils.CantFindMessage), errors.Is(err, utils.MessageNotLiked):
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
			return
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, utils.Ok)
}
