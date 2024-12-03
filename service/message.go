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

func GetAllComments() {

}

func DeleteComment() {

}

func SearchForComments() { //后期再添加的功能

}
