package dao

import (
	"fmt"
	"message-board/model"
)

func GetUserInfo(name string) (model.User, error) {
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
		return model.User{}, fmt.Errorf("wrong-username")
	}
}

func ChangeUserInfo() {

}

func DeleteUser() {

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
