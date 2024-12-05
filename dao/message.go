package dao

import (
	"message-board/model"
	"message-board/utils"
)

func GetAllComment() ([]model.Message, error) {
	var empty []model.Message
	var commentList []model.Message
	var comment model.Message
	query := "SELECT * FROM messages WHERE is_deleted = 0"
	rows, err := Db.Query(query)
	if err != nil {
		return empty, err
	}
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.IsDeleted, &comment.ParentID)
		if err != nil {
			return empty, err
		}
		commentList = append(commentList, comment)
	}
	if err = rows.Err(); err != nil {
		return empty, err
	}
	return commentList, nil
}

func AddComment(message model.Message) error {
	query := "SELECT * FROM users WHERE id=?"
	rows, err := Db.Query(query, message.UserID) //内部错误
	if err != nil {
		return err
	}
	if !rows.Next() { //参数错误
		return utils.InvalidID
	}
	query = "INSERT INTO messages (user_id,content) VALUES (?,?)"
	_, err = Db.Exec(query, message.UserID, message.Content)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(messageInfo model.Message) error {
	query := "UPDATE messages SET is_deleted=1 WHERE id = ?"
	result, err := Db.Exec(query, messageInfo.ID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return utils.CantFindMessage
	} else {
		return nil
	}
}
