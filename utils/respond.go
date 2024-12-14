package utils

import "errors"

type Response struct { //响应结构体
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var (
	Ok = Response{ //正常
		Status: 200,
		Info:   "OK",
	}
	WrongPwd = Response{ //密码错误
		Status: 400,
		Info:   "Wrong Password!",
	}
	InvalidID                 = errors.New("invalid userid")                                    //用户ID无效
	CantFindMessage           = errors.New("can't find this message")                           //找不到留言
	CantFindUser              = errors.New("can't find this user")                              //找不到用户
	WrongUsrName              = errors.New("wrong username")                                    //用户名错误
	InvalidUsername           = errors.New("invalid username")                                  //用户名无效
	MissingParam              = errors.New("more params needed")                                //缺少参数
	MissingToken              = errors.New("missing token")                                     //缺少token
	InvalidTokenSingingMethod = errors.New("invalid signing method")                            //jwt token签名方法无效
	InvalidToken              = errors.New("invalid token")                                     //无效token
	InvalidClaims             = errors.New("invalid claims")                                    //无效声明
	NotLoggedIn               = errors.New("not logged in")                                     //未登录
	ErrUnauthorized           = errors.New("unauthorized: only admins can perform this action") //未授权：只有管理员可以执行此操作
	SameInfoAsBefore          = errors.New("the new info is the same as before")                //要修改的新信息与原信息相同
	SamePassword              = errors.New("the new password is the same as before")            //新密码与原密码相同
	WrongOldPassword          = errors.New("wrong old password")                                //旧密码错误
	MessageAlreadyLiked       = errors.New("you've already liked this message")                 //已经点赞过了
	MessageNotLiked           = errors.New("you haven't liked this message")                    //没有点赞过
)

func ServerError(err error) Response { //服务器错误
	return Response{
		Status: 500,         //状态码500
		Info:   err.Error(), //错误信息
	}
}

func ClientError(err error) Response { //客户端错误
	return Response{
		Status: 400,
		Info:   err.Error(),
	}
}

func NotFoundError(err error) Response { //找不到内容
	return Response{
		Status: 404,
		Info:   err.Error(),
	}
}
