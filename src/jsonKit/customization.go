package jsonKit

type (
	MessageHook func(code string, message string, data interface{}) string

	ResponseHook func(resp *Response) any
)

// 根据 code、message、data，返回一个 新的message
var messageHook MessageHook

// 可用于修改响应对象属性的key值
var responseHook ResponseHook

func SetMsgHook(hook MessageHook) {
	messageHook = hook
}

func ClearMsgProcessor() {
	messageHook = nil
}

func SetRespHook(hook ResponseHook) {
	responseHook = hook
}

func ClearRespProcessor() {
	responseHook = nil
}
