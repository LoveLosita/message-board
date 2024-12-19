package api

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"message-board/model"
	"message-board/service"
	"message-board/utils"
	"strconv"
)

func SendMessage(ctx context.Context, c *app.RequestContext) {
	message := model.AdminGetMessage{}
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
	//fmt.Println(id)//测试
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

func AdminGetAllMessages(ctx context.Context, c *app.RequestContext) { //获取所有评论
	handlerID := int(c.GetFloat64("user_id"))                    //从上下文中获取用户的id
	allComments, err := service.AdminBuildMessageTree(handlerID) //调用service层的方法，获取所有评论
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrUnauthorized): //如果是权限错误
			c.JSON(consts.StatusUnauthorized, utils.ClientError(err)) //返回401
			return
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err)) //其他错误，返回500
			return
		}
	}
	combinedJson := map[string]interface{}{ //将所有评论和状态码组合成一个json
		"messages":     allComments,
		"respond code": utils.Ok,
	}
	c.JSON(consts.StatusOK, combinedJson) //返回所有评论和状态码
}

func UserGetAllMessages(ctx context.Context, c *app.RequestContext) { //获取所有评论
	allComments, err := service.UserBuildMessageTree() //调用service层的方法，获取所有评论
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

func DeleteMessage(ctx context.Context, c *app.RequestContext) { //管理员和评论发送人可使用,传入ID以删除评论
	messageID := c.Query("id")
	handlerID := int(c.GetFloat64("user_id"))
	if messageID == "" { //如果没有传入ID，那么返回错误，避免下方的转换出错
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.MissingParam))
		return
	}
	intMsgID, err := strconv.ParseInt(messageID, 10, 0)
	//fmt.Println("api", intMsgID)//测试
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
	err := c.BindJSON(&searchParams) //绑定前端传来的参数
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err)) //参数错误
		return
	}
	comments, err := service.SearchForMessages(searchParams.CommentID, searchParams.Content, searchParams.UserID,
		searchParams.Username) //调用service层的方法，进行查询
	if err != nil {
		switch {
		case errors.Is(err, utils.MissingParam), errors.Is(err, utils.InvalidID), errors.Is(err, utils.CantFindMessage): //参数不足
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

func ReplyMessage(ctx context.Context, c *app.RequestContext) {
	//首先读取前端传来的参数
	message := model.MessageReply{}
	err := c.BindJSON(&message)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	if message.Content == "" { //如果内容为空
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.MissingParam)) //返回错误
		return
	}
	//然后获取用户的id
	getID := c.GetFloat64("user_id") //从上下文中获取用户的id
	if getID == 0 {                  //id空白，表示未登录，不过一般在中间件就会被截止了
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.NotLoggedIn))
	}
	id := int(getID)
	//然后调用service层的方法
	message.UserID = id //将用户id赋值给message
	err = service.ReplyMessage(message)
	if err != nil {
		switch {
		case errors.Is(err, utils.InvalidID), errors.Is(err, utils.CantFindMessage): //如果是ID错误或者找不到留言
			c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err)) //其他错误，一般是服务器错误
		}
		return
	}
	//最后返回成功
	c.JSON(consts.StatusOK, utils.Ok)
}
