package base64Kit

import (
	"fmt"
	"testing"
)

func TestEncodeStringToStringAndDecodeStringToString(t *testing.T) {
	text := "test测试~！@#￥%……&*（）"

	base64Str := EncodeStringToString(text)
	fmt.Println("base64Str:", base64Str) // base64Str: dGVzdOa1i+ivlX7vvIFAI++/pSXigKbigKYmKu+8iO+8iQ==

	plainText, err := DecodeStringToString(base64Str)
	if err != nil {
		panic(err)
	}
	fmt.Println("plainText:", plainText) // plainText: test测试~！@#￥%……&*（）

	if text != plainText {
		panic("not equal")
	}
	fmt.Println("equal")
}
