package service

import (
	"message-board/dao"
	"message-board/model"
)

func SendComment(message model.Message) error {
	err := dao.AddComment(message)
	if err != nil {
		return err
	}
	return nil
}

func GetAllComments() ([]model.Message, error) {
	var commentList []model.Message
	commentList, err := dao.GetComment()
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func DeleteComment(message model.Message) error { //此处传结构体是为之后更新任意条件查找并删除做准备
	err := dao.DeleteComment(message)
	if err != nil {
		return err
	}
	return nil
}

func SearchForComments() { //后期再添加的功能

}
