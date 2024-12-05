package dao

import (
	"fmt"
	"message-board/model"
)

func GetComment() ([]model.Message, error) {
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
	fmt.Println(commentList)
	return commentList, nil
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
