package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonResplKit"
)

func main() {
	type resp struct {
		Code string      `json:"errorCode"`
		Msg  string      `json:"errorMessage"`
		Data interface{} `json:"result"`
	}
	var respProvider jsonResplKit.RespProvider = func(code, msg string, data interface{}) interface{} {
		msg = "[A] " + msg
		return &resp{
			Code: code,
			Msg:  msg,
			Data: data,
		}
	}

	jsonResplKit.MustSetUp(respProvider, []string{"chimera-lib/msg.properties"})
	// {"errorCode":"0","errorMessage":"[A] no error","result":1} <nil>
	fmt.Println(jsonResplKit.Seal("0", 1))
	// {"errorCode":"42","errorMessage":"[A] hello 测试","result":null} <nil>
	fmt.Println(jsonResplKit.Seal("42", nil, "测试"))
}
