package dao

import (
	"message-board/model"
	"message-board/utils"
)

func GetUserInfoByName(name string) (model.User, error) {
	user := model.User{}
	query := "SELECT id,nickname,username,password,created_at,updated_at,role FROM users WHERE username=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return model.User{}, err
	}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.NickName, &user.UserName, &user.PassWord, &user.CreatedAt, &user.UpdatedAt, &user.Role)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	} else {
		return model.User{}, utils.WrongUsrName
	}
}

func GetUserInfoByID(id int) (model.User, error) {
	user := model.User{}
	query := "SELECT id,nickname,username,password,created_at,updated_at,role FROM users WHERE id=?"
	rows, err := Db.Query(query, id)
	if err != nil {
		return model.User{}, err
	}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.NickName, &user.UserName, &user.PassWord, &user.CreatedAt, &user.UpdatedAt, &user.Role)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	} else {
		return model.User{}, utils.InvalidID
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
