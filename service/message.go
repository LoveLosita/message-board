package service

import (
	"message-board/auth"
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
)

func SendMessage(message model.AdminGetMessage) error { //发送评论
	err := dao.SendMessage(message) //调用dao层函数
	if err != nil {
		return err
	}
	return nil
}

func DeleteMessages(msgID int, handlerID int) error { //删除评论
	messages, err := dao.GetAllMessages() //获取该评论
	if err != nil {
		return err
	}
	//fmt.Println("sv", messages[0].UserID)//测试
	for _, message := range messages {
		if message.ID == msgID { //如果找到了这个评论
			if message.UserID == handlerID { //如果是评论的作者
				return dao.DeleteMessage(msgID) //调用dao层函数删除
			} else { //如果不是评论的作者
				return utils.ErrUnauthorized //返回错误
			}
		}
	}
	//下面是管理员删除评论的逻辑
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return err
	}
	if !result { //如果不是管理员
		return utils.ErrUnauthorized //返回错误
	}
	err = dao.DeleteMessage(msgID) //调用dao层函数
	if err != nil {
		return err
	}
	return nil
}

func SearchForMessages(commentID int, content string, userID int, username string) ([]model.AdminGetMessage, error) { //handlerID是管理员的id
	return dao.SearchForMessages(commentID, content, userID, username) //调用dao层函数
}

func LikeMessage(messageID int, handlerID int) error {
	result, err := dao.IfMessageExists(messageID)
	if err != nil {
		return err
	}
	if !result {
		return utils.CantFindMessage
	}
	result, err = dao.IfYouLikedThisMessage(messageID, handlerID)
	if err != nil {
		return err
	}
	if result { //如果已经点赞过了
		return utils.MessageAlreadyLiked //返回错误
	}
	return dao.LikeMessage(messageID, handlerID) //点赞
}

func DislikeMessage(messageID int, handlerID int) error {
	//首先要保证这个留言存在，然后要保证这个用户点赞过这个留言
	result, err := dao.IfMessageExists(messageID)
	if err != nil {
		return err
	}
	if !result {
		return utils.CantFindMessage
	}
	result, err = dao.IfYouLikedThisMessage(messageID, handlerID)
	if err != nil {
		return err
	}
	if !result { //如果没有点赞过
		return utils.MessageNotLiked //返回错误
	}
	return dao.DislikeMessage(messageID, handlerID) //取消点赞
}

func ReplyMessage(reply model.MessageReply) error { //回复评论
	err := dao.ReplyMessage(reply) //调用dao层函数
	if err != nil {
		return err
	}
	return nil
}

func AdminBuildMessageTree(handlerID int) ([]model.AdminGetMessage, error) { //构建评论树，用于前端展示
	result, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {
		return nil, err
	}
	if !result { //如果不是管理员
		return nil, utils.ErrUnauthorized //返回错误
	}
	adminMessages, err := dao.GetAllMessages() //获取所有评论
	if err != nil {
		return nil, err
	}
	for i := len(adminMessages) - 1; i >= 0; i-- { //从后往前遍历
		if adminMessages[i].ParentID != nil { //如果有父评论
			for j := 0; j < i; j++ { //寻找父评论
				if *adminMessages[i].ParentID == adminMessages[j].ID { //如果找到了父评论
					adminMessages[j].Replies = append(adminMessages[j].Replies, adminMessages[i]) //将子评论添加到父评论的Replies中
					break
				}
			}
		}
	}
	var resultList []model.AdminGetMessage       //定义一个管理员评论列表
	for _, adminMessage := range adminMessages { //遍历管理员评论
		if adminMessage.ParentID == nil { //如果没有父评论
			resultList = append(resultList, adminMessage) //将评论添加到管理员评论列表中
		}
	}
	return resultList, nil
}

func UserBuildMessageTree() ([]model.UserGetMessage, error) { //构建评论树，用于前端展示
	adminMessages, err := dao.GetAllMessages() //获取所有评论
	if err != nil {
		return nil, err
	}
	userMessages := AdminMessageToUserMessage(adminMessages) //将管理员评论转为用户评论
	for i := len(userMessages) - 1; i >= 0; i-- {            //从后往前遍历
		if userMessages[i].ParentID != nil { //如果有父评论
			for j := 0; j < i; j++ { //寻找父评论
				if *userMessages[i].ParentID == userMessages[j].ID { //如果找到了父评论
					userMessages[j].Replies = append(userMessages[j].Replies, userMessages[i]) //将子评论添加到父评论的Replies中
					break
				}
			}
		}
	}
	var resultList []model.UserGetMessage      //定义一个用户评论列表
	for _, userMessage := range userMessages { //遍历用户评论
		if userMessage.ParentID == nil { //如果没有父评论
			resultList = append(resultList, userMessage) //将评论添加到用户评论列表中
		}
	}
	return resultList, nil
}

func AdminMessageToUserMessage(adminMessages []model.AdminGetMessage) []model.UserGetMessage { //管理员评论转用户评论
	var userMessages []model.UserGetMessage      //定义一个用户评论列表
	for _, adminMessage := range adminMessages { //遍历管理员评论
		var userMessage model.UserGetMessage             //定义一个用户评论
		userMessage.ID = adminMessage.ID                 //将管理员评论的ID赋值给用户评论的ID
		userMessage.UserID = adminMessage.UserID         //将管理员评论的UserID赋值给用户评论的UserID
		userMessage.Content = adminMessage.Content       //将管理员评论的Content赋值给用户评论的Content
		userMessage.UpdatedAt = adminMessage.UpdatedAt   //将管理员评论的UpdatedAt赋值给用户评论的UpdatedAt
		userMessage.ParentID = adminMessage.ParentID     //将管理员评论的ParentID赋值给用户评论的ParentID
		userMessage.Likes = adminMessage.Likes           //将管理员评论的Likes赋值给用户评论的Likes
		userMessage.Replies = []model.UserGetMessage{}   //初始化用户评论的Replies
		userMessages = append(userMessages, userMessage) //将用户评论添加到用户评论列表中
	}
	return userMessages
}
