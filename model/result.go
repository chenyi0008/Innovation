package model

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Succeed(msg string, data interface{}) *Response {
	return &Response{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
}

func Failed(msg string, data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}
