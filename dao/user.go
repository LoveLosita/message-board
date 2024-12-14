package dao

import (
	"message-board/model"
	"message-board/utils"
)

func GetUserInfoByName(name string) (model.DisplayUser, error) {
	user := model.DisplayUser{}
	query := "SELECT id,nickname,username,created_at,updated_at,role FROM users WHERE username=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return model.DisplayUser{}, err
	}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.NickName, &user.UserName, &user.CreatedAt, &user.UpdatedAt, &user.Role)
		if err != nil {
			return model.DisplayUser{}, err
		}
		return user, nil
	} else {
		return model.DisplayUser{}, utils.WrongUsrName
	}
}

func GetUserInfoByID(id int) (model.DisplayUser, error) {
	user := model.DisplayUser{}
	query := "SELECT id,nickname,username,created_at,updated_at,role FROM users WHERE id=?"
	rows, err := Db.Query(query, id)
	if err != nil {
		return model.DisplayUser{}, err
	}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.NickName, &user.UserName, &user.CreatedAt, &user.UpdatedAt, &user.Role)
		if err != nil {
			return model.DisplayUser{}, err
		}
		return user, nil
	} else {
		return model.DisplayUser{}, utils.InvalidID
	}
}

func ChangeUserInfo(user model.NewUser) error {
	query := "UPDATE users SET nickname=?, username=?,password=?, role=? WHERE id=?"
	result, err := Db.Exec(query, user.NickName, user.UserName, user.PassWord, user.Role, user.TargetID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return utils.CantFindUser
	} else {
		return nil
	}
}

func DeleteUser(userID int) error {
	query := "DELETE FROM users WHERE id=?"
	result, err := Db.Exec(query, userID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return utils.CantFindUser
	}
	return nil
}

func AddUser(user model.User) error {
	query := "INSERT INTO users (nickname,username,password,role) VALUES (?,?,?,?)"
	_, err := Db.Exec(query, user.NickName, user.UserName, user.PassWord, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func IfUsernameExists(name string) (bool, error) {
	query := "SELECT username FROM users WHERE username=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return true, err
	}
	return rows.Next(), nil
}

func GetUserPasswordByName(name string) (int, string, error) {
	var password string
	var id int
	query := "SELECT id,password FROM users WHERE username=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return 0, "", err
	}
	if rows.Next() {
		err = rows.Scan(&id, &password)
		if err != nil {
			return 0, "", err
		}
		return id, password, nil
	} else {
		return 0, "", utils.InvalidUsername
	}
}

func ChangePassword(username string, newPwd string) error {
	query := "UPDATE users SET password=? WHERE username=?"
	result, err := Db.Exec(query, newPwd, username)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected() //返回受影响的行数
	if err != nil {
		return err
	}
	if affected == 0 { //如果受影响的行数为0，说明没有找到对应的用户
		return utils.CantFindUser
	} else {
		return nil
	}
}
