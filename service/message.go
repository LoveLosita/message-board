package service

import (
	"message-board/auth"
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
)

func SendMessage(message model.Message) error { //发送评论
	err := dao.AddMessage(message) //调用dao层函数
	if err != nil {
		return err
	}
	return nil
}

func GetAllMessages(handlerID int) ([]model.Message, error) { //获取所有评论
	var empty []model.Message
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return empty, err
	}
	if !result {
		return empty, utils.ErrUnauthorized //不是管理员，返回错误
	}
	var commentList []model.Message         //定义一个空的评论列表
	commentList, err = dao.GetAllMessages() //调用dao层函数
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func DeleteMessages(msgID int, handlerID int) error { //删除评论
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return err
	}
	if !result {
		return utils.ErrUnauthorized //不是管理员，返回错误
	}
	err = dao.DeleteMessage(msgID) //调用dao层函数
	if err != nil {
		return err
	}
	return nil
}

func SearchForMessages(commentID int, content string, userID int, username string, handlerID int) ([]model.Message, error) { //handlerID是管理员的id
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, utils.ErrUnauthorized //不是管理员，返回错误
	}
	return dao.SearchForMessages(commentID, content, userID, username) //是管理员，调用dao层函数
}

func LikeMessage(messageID int, handlerID int) error {
	result, err := dao.IfMessageExists(messageID)
	if err != nil {
		return err
	}
	if !result {
		return utils.CantFindMessage
	}
	result, err = dao.IfYouLikedThisMessage(messageID, handlerID)
	if err != nil {
		return err
	}
	if result { //如果已经点赞过了
		return utils.MessageAlreadyLiked //返回错误
	}
	return dao.LikeMessage(messageID, handlerID) //点赞
}

func DislikeMessage(messageID int, handlerID int) error {
	//首先要保证这个留言存在，然后要保证这个用户点赞过这个留言
	result, err := dao.IfMessageExists(messageID)
	if err != nil {
		return err
	}
	if !result {
		return utils.CantFindMessage
	}
	result, err = dao.IfYouLikedThisMessage(messageID, handlerID)
	if err != nil {
		return err
	}
	if !result { //如果没有点赞过
		return utils.MessageNotLiked //返回错误
	}
	return dao.DislikeMessage(messageID, handlerID) //取消点赞
}
