package aesCbcPkcs7Kit

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	text := "test测试~！@#￥%……&*（）——+-="
	key := "0123456789abcdef"
	iv := "0123456789abcdef"

	cipherText, err := EncryptToBase64([]byte(text), []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(cipherText)

	plainTextData, err := DecryptFromBase64(cipherText, []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plainTextData))

	if text != string(plainTextData) {
		panic("not equal")
	}
}

func TestHex(t *testing.T) {
	text := "test测试~！@#￥%……&*（）——+-="
	key := "0123456789abcdef"
	iv := "0123456789abcdef"

	cipherText, err := EncryptToHex([]byte(text), []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(cipherText)

	plainTextData, err := DecryptFromHex(cipherText, []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plainTextData))

	if text != string(plainTextData) {
		panic("not equal")
	}
}
