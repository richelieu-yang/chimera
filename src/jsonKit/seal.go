package jsonKit

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/msgKit"
)

type (
	// Response 响应给前端的json对象.
	Response struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

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
