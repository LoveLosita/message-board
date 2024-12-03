package service

import (
	"fmt"
	"message-board/dao"
	"message-board/model"
)

func UserLogin(username string, password string) (bool, error) {
	daoUser, err := dao.GetUserInfo(username)
	if err != nil {
		return false, err
	}
	if password == daoUser.PassWord {
		return true, nil
	} else {
		return false, nil
	}
}

func UserRegister(user model.User) error {
	if user.UserName == "" || user.NickName == "" || user.Role == "" || user.PassWord == "" {
		return fmt.Errorf("more-info-needed")
	}
	ifExist, err := dao.IfUsernameExists(user.UserName)
	if ifExist {
		return fmt.Errorf("invalid-username")
	}
	if err != nil {
		return err
	}
	err = dao.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func ShowUserInfo() { //后期再添加的功能

}

func ChangeUserInfo() { //后期再添加的功能

}

func DeleteUser() { //后期再添加的功能

}
