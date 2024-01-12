package base64Kit

import (
	"fmt"
	"testing"
)

func TestEncodeStringToStringAndDecodeStringToString(t *testing.T) {
	text := "test测试~!@#$%^&*()"

	base64Str := EncodeStringToString(text)
	fmt.Println("base64Str:", base64Str)

	plainText, err := DecodeStringToString(base64Str)
	if err != nil {
		panic(err)
	}
	fmt.Println("plainText:", plainText)

	if text != plainText {
		panic("not equal")
	}
	fmt.Println("equal")
}
