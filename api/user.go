package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {

}

func UserRegister(ctx context.Context, c *app.RequestContext) {

}

//以下是管理员专属功能

func ShowUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}

func DeleteUser(ctx context.Context, c *app.RequestContext) { //后期再添加的功能

}
