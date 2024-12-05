package service

import (
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
)

func UserLogin(username string, password string) (bool, error) {
	daoUser, err := dao.GetUserInfo(username)
	if err != nil { //找不到用户名或者内部错误
		return false, err
	}
	if password == daoUser.PassWord {
		return true, nil
	} else {
		return false, nil //密码错误
	}
}

func UserRegister(user model.User) error {
	if user.UserName == "" || user.NickName == "" || user.Role == "" || user.PassWord == "" {
		return utils.MissingParam
	}
	ifExist, err := dao.IfUsernameExists(user.UserName)
	if ifExist {
		return utils.InvalidUsername
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
