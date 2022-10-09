package jsonKit

type (
	MsgProcessor func(code string, message string, data interface{}) string

	JsonResponseProcessor func(resp *JsonResponse) any
)

// 不为nil的情况，将根据 code、message、data 返回一个 新的message
var msgProcessor MsgProcessor

// 不为nil的情况，将根据 code 和 msg 返回一个新的 msg
var jsonResponseProcessor JsonResponseProcessor

// SetMsgProcessor
/*
可以通过此方法实现修改 JsonResponse.Message，可用于: 告知前端此响应是哪个服务返回的，便于后续定位问题.
*/
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
