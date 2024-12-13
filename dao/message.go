package dao

import (
	"database/sql"
	"message-board/model"
	"message-board/utils"
)

func GetAllComment() ([]model.Message, error) { //获取所有评论
	var empty []model.Message                              //空的返回值
	var commentList []model.Message                        //返回的评论列表
	var comment model.Message                              //单个评论
	query := "SELECT * FROM messages WHERE is_deleted = 0" //查询语句
	rows, err := Db.Query(query)                           //查询结果
	if err != nil {                                        //如果查询出错，那么返回空值和错误
		return empty, err
	}
	for rows.Next() { //遍历查询结果
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt,
			&comment.IsDeleted, &comment.ParentID) //将查询结果赋值给comment
		if err != nil { //如果遍历出错，那么返回空值和错误
			return empty, err
		}
		commentList = append(commentList, comment) //将查询结果添加到返回值中
	}
	if err = rows.Err(); err != nil { //如果遍历出错，那么返回空值和错误
		return empty, err
	}
	return commentList, nil //返回查询结果
}

func AddComment(message model.Message) error { //发送评论
	query := "SELECT * FROM users WHERE id=?"    //查询语句
	rows, err := Db.Query(query, message.UserID) //内部错误
	if err != nil {                              //如果查询出错，那么返回错误
		return err
	}
	if !rows.Next() { //参数错误
		return utils.InvalidID //返回ID错误
	}
	query = "INSERT INTO messages (user_id,content) VALUES (?,?)"
	_, err = Db.Exec(query, message.UserID, message.Content) //插入评论
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(msgID int) error {
	query := "UPDATE messages SET is_deleted=1 WHERE id = ?"
	result, err := Db.Exec(query, msgID)
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

func SearchForComments(commentID int, content string, userID int, username string) ([]model.Message, error) {
	var empty []model.Message       //空的返回值
	var commentList []model.Message //返回的评论列表
	var comment model.Message       //单个评论
	var query string                //查询语句
	var rows *sql.Rows              //查询结果
	var err error                   //错误

	if commentID != 0 { //首先第一个判断评论id是否为0，如果不为0，那么就是通过评论id查找
		query = "SELECT * FROM messages WHERE id = ? AND is_deleted = 0"
		rows, err = Db.Query(query, commentID)
	} else if content != "" { //如果评论id为0，那么就是通过评论内容查找
		query = "SELECT * FROM messages WHERE content LIKE ? AND is_deleted = 0"
		rows, err = Db.Query(query, "%"+content+"%")
	} else if userID != 0 { //如果评论内容为空，那么就是通过用户id查找
		query = "SELECT * FROM messages WHERE user_id = ? AND is_deleted = 0"
		rows, err = Db.Query(query, userID)
	} else if username != "" { //如果用户id为0，那么就是通过用户名查找
		query = `SELECT m.* FROM messages m JOIN users u ON m.user_id = u.id WHERE u.username = ? AND m.is_deleted = 0`
		rows, err = Db.Query(query, username)
	} else { //如果用户名为空，那么就是参数不足
		return empty, utils.MissingParam
	}
	if err != nil { //如果查询出错，那么返回空值和错误
		return empty, err
	}
	if !rows.Next() { //如果找不到评论，那么返回空值和错误
		return empty, utils.CantFindMessage
	}
	for rows.Next() { //遍历查询结果
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.IsDeleted, &comment.ParentID)
		if err != nil {
			return empty, err
		}
		commentList = append(commentList, comment) //将查询结果添加到返回值中
	}
	if err = rows.Err(); err != nil { //如果遍历出错，那么返回空值和错误
		return empty, err
	}
	return commentList, nil
}
