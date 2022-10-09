package jsonKit

type (
	MessageProcessor func(code string, message string, data interface{}) string

	ResponseProcessor func(resp *Response) any
)

// 根据 code、message、data 返回一个 新的message
var msgProcessor MessageProcessor

// 可用于修改响应对象属性的key值
var responseProcessor ResponseProcessor

func SetMsgProcessor(processor MessageProcessor) {
	msgProcessor = processor
}

func ClearMsgProcessor() {
	msgProcessor = nil
}

func SetRespProcessor(processor ResponseProcessor) {
	responseProcessor = processor
}

func ClearRespProcessor() {
	responseProcessor = nil
}
