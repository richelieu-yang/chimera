package hexKit

import (
	"fmt"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	text := "test测试~!@#$%^&*()"

	cipherData := Encode([]byte(text))
	fmt.Println("cipherTest:", string(cipherData))

	plainData, err := Decode(cipherData)
	if err != nil {
		panic(err)
	}
	fmt.Println("plainText:", string(plainData))

	if text != string(plainData) {
		panic("not equal")
	}
	fmt.Println("equal")
}
