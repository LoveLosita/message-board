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
	InvalidID       = errors.New("invalid userid")
	CantFindMessage = errors.New("can't find this message")
	WrongUsrName    = errors.New("wrong username")
	InvalidUsername = errors.New("invalid username")
	MissingParam    = errors.New("more params needed")
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
