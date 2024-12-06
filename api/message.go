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
)

func SendComment(ctx context.Context, c *app.RequestContext) {
	message := model.Message{}
	err := c.BindJSON(&message)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	getID := c.GetFloat64("user_id") //从上下文中获取用户的id
	if getID == 0 {                  //id空白，表示未登录，不过一般在中间件就会被截止了
		c.JSON(consts.StatusBadRequest, utils.ClientError(utils.NotLoggedIn))
	}
	id := int(getID)
	fmt.Println(id)
	message.UserID = id
	err = service.SendComment(message)
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

func GetAllComments(ctx context.Context, c *app.RequestContext) {
	allComments, err := service.GetAllComments()
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		return
	}
	combinedJson := map[string]interface{}{
		"messages":     allComments,
		"respond code": utils.Ok,
	}
	c.JSON(consts.StatusOK, combinedJson)
}

func DeleteComment(ctx context.Context, c *app.RequestContext) {
	messageToDelete := model.Message{}
	err := c.BindJSON(&messageToDelete)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.ClientError(err))
		return
	}
	err = service.DeleteComment(messageToDelete)
	if err != nil {
		switch {
		case errors.Is(err, utils.CantFindMessage):
			c.JSON(consts.StatusBadRequest, utils.NotFoundError(err))
		default:
			c.JSON(consts.StatusInternalServerError, utils.ServerError(err))
		}
		return
	}
	c.JSON(consts.StatusOK, utils.Ok)
}

func SearchForComments(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}
