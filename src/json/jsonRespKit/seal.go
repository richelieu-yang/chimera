package jsonRespKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// Seal
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func Seal(code string, data interface{}, msgArgs ...interface{}) (string, error) {
	return SealFully(code, "", data, msgArgs...)
}

// SealFully
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func SealFully(code, msg string, data interface{}, msgArgs ...interface{}) (string, error) {
	if strKit.IsEmpty(msg) {
		msg = msgMap[code]
	}
	if strKit.IsNotEmpty(msg) && msgArgs != nil {
		msg = fmt.Sprintf(msg, msgArgs...)
	}

	bean := provider(code, msg, data)
	return api.MarshalToString(bean)
}
