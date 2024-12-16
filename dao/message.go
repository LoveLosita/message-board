package dao

import (
	"database/sql"
	"message-board/model"
	"message-board/utils"
)

func GetAllMessages() ([]model.Message, error) { //获取所有留言
	var empty []model.Message                                                  //空的返回值
	var MessageList []model.Message                                            //返回的评论列表 	//单个评论
	query := "SELECT * FROM messages WHERE is_deleted = 0 ORDER BY created_at" //查询语句
	rows, err := Db.Query(query)                                               //查询结果
	if err != nil {                                                            //如果查询出错，那么返回空值和错误
		return empty, err
	}
	for rows.Next() { //遍历查询结果
		var message model.Message // 每次循环创建一个新的实例
		err = rows.Scan(&message.ID, &message.UserID, &message.Content, &message.CreatedAt, &message.UpdatedAt,
			&message.IsDeleted, &message.ParentID, &message.Likes) //将查询结果赋值给comment
		if err != nil { //如果遍历出错，那么返回空值和错误
			return empty, err
		}
		MessageList = append(MessageList, message) //将查询结果添加到返回值中
	}
	if err = rows.Err(); err != nil { //如果遍历出错，那么返回空值和错误
		return empty, err
	}
	return MessageList, nil //返回查询结果
}

func SendMessage(message model.Message) error { //发送留言
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

func DeleteMessage(msgID int) error { //删除留言
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

func SearchForMessages(commentID int, content string, userID int, username string) ([]model.Message, error) { //搜索评论
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
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.IsDeleted, &comment.ParentID, &comment.Likes) //将查询结果赋值给comment
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

func LikeMessage(messageID int, userID int) error { //点赞
	query := "INSERT INTO likes (beliked_message_id,like_user_id) VALUES (?,?)" //插入点赞
	_, err := Db.Exec(query, messageID, userID)                                 //执行插入
	if err != nil {
		return err
	}
	query = "UPDATE messages SET `like` = `like` + 1 WHERE id = ?"
	result, err := Db.Exec(query, messageID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected() //返回受影响的行数
	if err != nil {
		return err
	}
	if affected == 0 {
		return utils.CantFindMessage //找不到评论
	}
	return nil
}

func IfYouLikedThisMessage(messageID int, userID int) (bool, error) { //检查是否点赞
	query := "SELECT * FROM likes WHERE beliked_message_id = ? AND like_user_id = ?" //查询语句
	rows, err := Db.Query(query, messageID, userID)
	if err != nil {
		return true, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, err
}

func IfMessageExists(messageID int) (bool, error) { //检查留言是否存在
	query := "SELECT * FROM messages WHERE id = ?"
	rows, err := Db.Query(query, messageID)
	if err != nil {
		return true, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, err
}

func DislikeMessage(messageID int, userID int) error { //取消点赞
	//第一步，删除点赞记录
	query := "DELETE FROM likes WHERE beliked_message_id = ? AND like_user_id = ?"
	result, err := Db.Exec(query, messageID, userID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return utils.CantFindMessage
	}
	//第二步，更新点赞数
	query = "UPDATE messages SET `like` = `like` - 1 WHERE id = ?"
	result, err = Db.Exec(query, messageID)
	if err != nil {
		return err
	}
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return utils.CantFindMessage
	}
	return nil
}

func ReplyMessage(message model.ReplyMessage) error {
	//首先要保证这个用户存在
	query := "SELECT * FROM users WHERE id=?"    //查询是否有这个用户
	rows, err := Db.Query(query, message.UserID) //内部错误
	if err != nil {                              //如果查询出错，那么返回错误
		return err
	}
	if !rows.Next() { //用户不存在
		return utils.InvalidID //返回ID错误
	}
	//然后要保证要回复的留言存在
	query = "SELECT * FROM messages WHERE id=?" //查询是否有这个留言
	rows, err = Db.Query(query, message.ParentID)
	if err != nil {
		return err
	}
	if !rows.Next() { //留言不存在
		return utils.CantFindMessage //返回留言错误
	}
	//最后插入回复
	query = "INSERT INTO messages (user_id,content,parent_id) VALUES (?,?,?)"  //插入留言语句
	_, err = Db.Exec(query, message.UserID, message.Content, message.ParentID) //执行插入
	if err != nil {
		return err
	}
	return nil
}
