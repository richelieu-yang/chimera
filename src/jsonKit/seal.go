package jsonKit

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/msgKit"
)

// Seal
/*
@param args 不传参的情况，值为nil
*/
func Seal(code string, msgArgs ...interface{}) string {
	return SealWithData(code, nil, msgArgs...)
}

func SealWithData(code string, data interface{}, msgArgs ...interface{}) string {
	json, _ := SealFully(jsoniter.ConfigDefault, code, "", data, msgArgs...)
	return json
}

func SealFully(api jsoniter.API, code, message string, data interface{}, msgArgs ...interface{}) (json string, err error) {
	if api == nil {
		api = jsoniter.ConfigDefault
	}

	message = getFinalMessage(code, message, data, msgArgs...)
	resp := &Response{Code: code, Message: message, Data: data}

	if responseHook == nil {
		return MarshalToStringWithJsoniterApi(api, resp)
	}
	return MarshalToStringWithJsoniterApi(api, responseHook(resp))
}

func getFinalMessage(code, msg string, data interface{}, msgArgs ...interface{}) string {
	if strKit.IsEmpty(msg) {
		msg = msgKit.GetMsg(code)
	}
	if strKit.IsNotEmpty(msg) && msgArgs != nil {
		msg = fmt.Sprintf(msg, msgArgs...)
	}
	if messageHook != nil {
		msg = messageHook(code, msg, data)
	}
	return msg
}
