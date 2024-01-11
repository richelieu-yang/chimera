package jsonRespKit

type (
	// RespProvider
	/*
		@return (1) 返回值是一个结构体实例指针
				(2) 结构体建议加上json tag
	*/
	RespProvider func(code, msg string, data interface{}) interface{}

	// respBean 默认的响应结构体
	respBean struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

var provider RespProvider = func(code, msg string, data interface{}) interface{} {
	return &respBean{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
