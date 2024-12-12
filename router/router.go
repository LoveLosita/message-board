package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"message-board/api"
	"message-board/middleware"
)

func RegisterRouters() {
	h := server.Default()

	userGroup := h.Group("/user") //创建用户组

	userGroup.POST("/login", api.UserLogin)       //user->login //checked
	userGroup.POST("/register", api.UserRegister) //user->register //checked

	messageGroup := h.Group("/message") //创建留言组

	messageGroup.POST("/submit", middleware.JWTAuthMiddleware(), api.SendComment) //message->submit //checked

	adminGroup := h.Group("/admin")                 //创建管理组
	messageSubGroup := adminGroup.Group("/message") //创建管理->留言组

	userSubGroup := adminGroup.Group("/user") //创建管理->用户组

	messageSubGroup.GET("/get-all", middleware.JWTAuthMiddleware(), api.GetAllComments)   //admin->message->get-all //checked
	messageSubGroup.DELETE("/delete", middleware.JWTAuthMiddleware(), api.DeleteComment)  //admin->message->delete //checked
	userSubGroup.GET("/show", middleware.JWTAuthMiddleware(), api.ShowUserInfo)           //admin->user->show //checked
	userSubGroup.POST("/change", middleware.JWTAuthMiddleware(), api.ChangeUserInfo)      //admin->user->change //checked
	userSubGroup.DELETE("/delete", middleware.JWTAuthMiddleware(), api.DeleteUser)        //admin->user->delete //checked
	messageSubGroup.GET("/search", middleware.JWTAuthMiddleware(), api.SearchForComments) //admin->message->search //checked
	h.Spin()

}
