package api

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterRouters() {
	h := server.Default()
	h.POST("/user/login", UserLogin)
	h.POST("/user/register", UserRegister) //验收通过
	h.POST("/admin/message/submit", SendComment)
	h.GET("/admin/message/get-all", GetAllComments)
	h.DELETE("/admin/message/delete", DeleteComment)
	h.GET("/admin/user/show", ShowUserInfo)            //后期开发的功能
	h.POST("/admin/user/change", ChangeUserInfo)       //后期开发的功能
	h.DELETE("/admin/user/delete", DeleteUser)         //后期开发的功能
	h.POST("/admin/message/search", SearchForComments) //后期开发的功能
	h.Spin()
}
