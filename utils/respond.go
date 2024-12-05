package utils

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
		Status: 404,
		Info:   "Wrong Password!",
	}
)

func CustomError(err error) Response {
	return Response{
		Status: 404,
		Info:   err.Error(),
	}
}

func CustomSuccess(message string) Response {
	return Response{
		Status: 200,
		Info:   message,
	}
}
