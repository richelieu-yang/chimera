package jsonKit

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/msgKit"
)

// Seal
/*
@param args 不传参的情况，值为nil
*/
func Seal(code string, args ...interface{}) string {
	return SealWithData(code, nil, args...)
}

func SealWithData(code string, data interface{}, args ...interface{}) string {
	json, _ := SealFully(jsoniter.ConfigDefault, code, "", data, args...)
	return json
}

func SealFully(api jsoniter.API, code, message string, data interface{}, args ...interface{}) (json string, err error) {
	if api == nil {
		api = jsoniter.ConfigDefault
	}

	message = getFinalMessage(code, message, data, args...)
	resp := &Response{Code: code, Message: message, Data: data}

	var obj any
	if responseProcessor != nil {
		obj = responseProcessor(resp)
	} else {
		obj = resp
	}
	return MarshalToStringWithJsoniterApi(api, obj)
}

func getFinalMessage(code, msg string, data interface{}, args ...interface{}) string {
	if strKit.IsEmpty(msg) {
		msg = msgKit.GetMsg(code)
	}
	if strKit.IsNotEmpty(msg) && args != nil {
		msg = fmt.Sprintf(msg, args...)
	}
	if msgProcessor != nil {
		msg = msgProcessor(code, msg, data)
	}
	return msg
}
