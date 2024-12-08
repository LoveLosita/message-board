package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"message-board/api"
	"message-board/middleware"
)

func RegisterRouters() {
	h := server.Default()

	userGroup := h.Group("/user") //创建用户组

	userGroup.POST("/login", api.UserLogin)
	userGroup.POST("/register", api.UserRegister) //验收通过

	messageGroup := h.Group("/message") //创建留言组

	messageGroup.POST("/submit", middleware.JWTAuthMiddleware(), api.SendComment) //message->submit

	adminGroup := h.Group("/admin")                 //创建管理组
	messageSubGroup := adminGroup.Group("/message") //创建管理->留言组

	userSubGroup := adminGroup.Group("/user") //创建管理->用户组

	messageSubGroup.GET("/get-all", api.GetAllComments)    //admin->message->get-all
	messageSubGroup.DELETE("/delete", api.DeleteComment)   //admin->message->delete
	userSubGroup.GET("/show", api.ShowUserInfo)            //admin->user->show
	userSubGroup.POST("/change", api.ChangeUserInfo)       //admin->user->change
	userSubGroup.DELETE("/delete", api.DeleteUser)         //admin->user->delete
	messageSubGroup.POST("/search", api.SearchForComments) //admin->message->search
	h.Spin()

}
