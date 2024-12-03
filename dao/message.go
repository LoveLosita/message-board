package dao

import "message-board/model"

func GetComment() {

}

func AddComment(message model.Message) error {
	query := "INSERT INTO messages (user_id,content) VALUES (?,?)"
	_, err := Db.Exec(query, message.UserID, message.Content)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment() {

}
