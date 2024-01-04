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
	fmt.Println(cipherText) // 8IGfmJx9UzhZAbdS2Jqqh5YlwcjISpDdm429JVKnK13/bafYcNNnRux4OolbmKxS

	plainTextData, err := DecryptFromBase64(cipherText, []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plainTextData)) // test测试~！@#￥%……&*（）——+-=

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
	fmt.Println(cipherText) // f0819f989c7d53385901b752d89aaa879625c1c8c84a90dd9b8dbd2552a72b5dff6da7d870d36746ec783a895b98ac52

	plainTextData, err := DecryptFromHex(cipherText, []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plainTextData)) // test测试~！@#￥%……&*（）——+-=

	if text != string(plainTextData) {
		panic("not equal")
	}
}
