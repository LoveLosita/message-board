package utils

import "errors"

type Response struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var (
	Ok = Response{
		Status: 200,
		Info:   "OK",
	}
	WrongPwd = Response{
		Status: 400,
		Info:   "Wrong Password!",
	}
	InvalidID                 = errors.New("invalid userid")
	CantFindMessage           = errors.New("can't find this message")
	CantFindUser              = errors.New("can't find this user")
	WrongUsrName              = errors.New("wrong username")
	InvalidUsername           = errors.New("invalid username")
	MissingParam              = errors.New("more params needed")
	MissingToken              = errors.New("missing token")
	InvalidTokenSingingMethod = errors.New("invalid signing method")
	InvalidToken              = errors.New("invalid token")
	InvalidClaims             = errors.New("invalid claims")
	NotLoggedIn               = errors.New("not logged in")
	ErrUnauthorized           = errors.New("unauthorized: only admins can perform this action")
)

func ServerError(err error) Response {
	return Response{
		Status: 500,
		Info:   err.Error(),
	}
}

func ClientError(err error) Response {
	return Response{
		Status: 400,
		Info:   err.Error(),
	}
}

func NotFoundError(err error) Response {
	return Response{
		Status: 404,
		Info:   err.Error(),
	}
}
