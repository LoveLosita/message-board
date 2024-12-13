package service

import (
	"message-board/auth"
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
)

func SendComment(message model.Message) error { //发送评论
	err := dao.AddMessage(message) //调用dao层函数
	if err != nil {
		return err
	}
	return nil
}

func GetAllComments(handlerID int) ([]model.Message, error) { //获取所有评论
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

func DeleteComment(msgID int, handlerID int) error { //删除评论
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

func SearchForComments(commentID int, content string, userID int, username string, handlerID int) ([]model.Message, error) { //handlerID是管理员的id
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, utils.ErrUnauthorized //不是管理员，返回错误
	}
	return dao.SearchForMessages(commentID, content, userID, username) //是管理员，调用dao层函数
}
