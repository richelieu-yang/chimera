package hexKit

import (
	"fmt"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	text := "test测试~!@#$%^&*()"

	cipherData := Encode([]byte(text))
	fmt.Println("cipherText:", string(cipherData)) // cipherText: 74657374e6b58be8af957e21402324255e262a2829

	plainData, err := Decode(cipherData) // plainText: test测试~!@#$%^&*()
	if err != nil {
		panic(err)
	}
	fmt.Println("plainText:", string(plainData))

	if text != string(plainData) {
		panic("not equal")
	}
	fmt.Println("equal")
}

func TestEncodeToStringAndDecodeToString(t *testing.T) {
	text := "test测试~!@#$%^&*()"

	cipherText := EncodeToString([]byte(text))
	fmt.Println("cipherText:", cipherText) // cipherText: 74657374e6b58be8af957e21402324255e262a2829

	plainData, err := DecodeString(cipherText)
	if err != nil {
		panic(err)
	}
	fmt.Println("plainText:", string(plainData)) // plainText: test测试~!@#$%^&*()

	if text != string(plainData) {
		panic("not equal")
	}
	fmt.Println("equal")
}
