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

func DeleteComment() {

}

func SearchForComments() { //后期再添加的功能

}
