package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"message-board/api"
	"message-board/middleware"
)

func RegisterRouters() {
	h := server.Default()

	userGroup := h.Group("/user") //创建用户组

	userGroup.POST("/login", api.UserLogin)                //user->login //checked
	userGroup.POST("/register", api.UserRegister)          //user->register //checked
	userGroup.POST("/change_password", api.ChangePassword) //user->change_password

	messageGroup := h.Group("/message") //创建留言组

	messageGroup.POST("/submit", middleware.JWTAuthMiddleware(), api.SendMessage)       //message->submit //checked
	messageGroup.POST("/like", middleware.JWTAuthMiddleware(), api.LikeMessage)         //message->like //checked
	messageGroup.DELETE("/dislike", middleware.JWTAuthMiddleware(), api.DislikeMessage) //message->dislike //checked
	messageGroup.GET("/get-all", api.UserGetAllMessages)                                //message->get-all //checked
	messageGroup.GET("/search", api.SearchForMessages)                                  //message->search //checked
	messageGroup.POST("/reply", middleware.JWTAuthMiddleware(), api.ReplyMessage)       //message->reply

	adminGroup := h.Group("/admin")                 //创建管理组
	messageSubGroup := adminGroup.Group("/message") //创建管理->留言组

	userSubGroup := adminGroup.Group("/user") //创建管理->用户组

	messageSubGroup.DELETE("/delete", middleware.JWTAuthMiddleware(), api.DeleteMessage)     //admin->message->delete //checked
	userSubGroup.GET("/show", middleware.JWTAuthMiddleware(), api.ShowUserInfo)              //admin->user->show //checked
	userSubGroup.POST("/change", middleware.JWTAuthMiddleware(), api.ChangeUserInfo)         //admin->user->change //checked
	userSubGroup.DELETE("/delete", middleware.JWTAuthMiddleware(), api.DeleteUser)           //admin->user->delete //checked
	messageSubGroup.GET("/get-all", middleware.JWTAuthMiddleware(), api.AdminGetAllMessages) //admin->message->get-all

	h.Spin()

}
