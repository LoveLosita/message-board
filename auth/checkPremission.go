package auth

import (
	"message-board/dao"
	"message-board/utils"
)

func CheckPermission(handlerID int) (bool, error) { // 检查用户权限
	user, err := dao.GetUserInfoByID(handlerID) //通过handlerID获取用户信息
	if err != nil {
		return false, err
	}
	if user.Role != "admin" { //判断是否是管理员
		return false, utils.ErrUnauthorized //不是管理员，返回错误
	}
	return true, nil
}
