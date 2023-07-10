package jsonResplKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/msgKit"
)

func Seal(code string, data interface{}, msgArgs ...interface{}) string {
	return SealFully(code, "", data, msgArgs...)
}

// SealFully
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func SealFully(code, msg string, data interface{}, msgArgs ...interface{}) string {
	if strKit.IsEmpty(msg) {
		msg = msgKit.GetMsg(code)
	}
	if strKit.IsNotEmpty(msg) && msgArgs != nil {
		msg = fmt.Sprintf(msg, msgArgs...)
	}
	if msgProcessor != nil {
		msg = msgProcessor(msg)
	}

	resp := provider(code, msg, data)
	// 忽略错误
	json, _ := api.MarshalToString(resp)
	return json
}
