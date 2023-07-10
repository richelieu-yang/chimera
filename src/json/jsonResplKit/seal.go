package jsonResplKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/msgKit"
)

func Seal(code string, data interface{}, msgArgs ...interface{}) string {
	return SealFully(code, "", data, msgArgs...)
}

func SealFully(code, msg string, data interface{}, msgArgs ...interface{}) string {
	if strKit.IsEmpty(msg) {
		msg = msgKit.GetMsg(code)
	}
	if strKit.IsNotEmpty(msg) && msgArgs != nil {
		msg = fmt.Sprintf(msg, msgArgs...)
	}
	// 供外部对最终message进行二开
	if msgProcessor != nil {
		msg = msgProcessor(msg)
	}

	resp := provider(code, msg, data)
	// 忽略错误
	json, _ := api.MarshalToString(resp)
	return json
}
