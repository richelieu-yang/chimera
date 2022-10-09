package jsonKit

type (
	MsgProcessor func(code string, message string, data interface{}) string

	JsonResponseProcessor func(resp *JsonResponse) any
)

// 根据 code、message、data 返回一个 新的message
var msgProcessor MsgProcessor

// 可用于修改响应对象属性的key值
var jsonResponseProcessor JsonResponseProcessor

func SetMsgProcessor(processor MsgProcessor) {
	msgProcessor = processor
}

func ClearMsgProcessor() {
	msgProcessor = nil
}

func SetJsonResponseProcessor(processor JsonResponseProcessor) {
	jsonResponseProcessor = processor
}

func ClearJsonResponseProcessor() {
	jsonResponseProcessor = nil
}
