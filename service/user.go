package service

import (
	"message-board/auth"
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
)

func UserLogin(username string, password string) (int, bool, error) {
	daoUser, err := dao.GetUserInfoByName(username)
	if err != nil { //找不到用户名或者内部错误
		return 0, false, err
	}
	if password == daoUser.PassWord {
		return daoUser.ID, true, nil
	} else {
		return 0, false, nil //密码错误
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

func ShowUserInfo(userName string, userID int, handlerID int) (model.User, error) { //用户名和用户ID二选一，优先使用用户名
	var emptyUser model.User
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return emptyUser, err
	}
	if !result {
		return emptyUser, utils.ErrUnauthorized //不是管理员，返回错误
	}
	if userName != "" {
		user, err := dao.GetUserInfoByName(userName)
		if err != nil {
			return emptyUser, err
		}
		return user, nil
	} else {
		user, err := dao.GetUserInfoByID(userID)
		if err != nil {
			return emptyUser, err
		}
		return user, nil
	}
}

func ChangeUserInfo(handlerID int, user model.NewUser) error {
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return err
	}
	if !result {
		return utils.ErrUnauthorized //不是管理员，返回错误
	}
	currentUser, err := dao.GetUserInfoByID(user.TargetID)
	if err != nil {
		return err
	}
	if currentUser.UserName == user.UserName || currentUser.NickName == user.NickName || //如果新信息和旧信息相同，那么就不修改
		currentUser.PassWord == user.PassWord || currentUser.Role == user.Role {
		return utils.SameInfoAsBefore
	}
	if user.UserName == "" { //如果用户名为空，那么就不修改用户名
		user.UserName = currentUser.UserName
	}
	if user.NickName == "" { //如果昵称为空，那么就不修改昵称
		user.NickName = currentUser.NickName
	}
	if user.PassWord == "" { //如果密码为空，那么就不修改密码
		user.PassWord = currentUser.PassWord
	}
	if user.Role == "" { //如果角色为空，那么就不修改角色
		user.Role = currentUser.Role
	}
	err = dao.ChangeUserInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID int, handlerID int) error { //后期再添加的功能
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return err
	}
	if !result {
		return utils.ErrUnauthorized //不是管理员，返回错误
	}
	err = dao.DeleteUser(userID) //调用dao层的方法，删除用户
	if err != nil {
		return err
	}
	return nil
}
