package aesKit

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesCbcPkcs7EncryptAndDecrypt(t *testing.T) {
	// 测试数据
	data := []byte("Hello, world!")
	key := []byte("00000yozo_config")
	iv := []byte("00000yozo_config")

	// 加密
	crypted, err := AesCbcPkcs7Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后的数据：", base64.StdEncoding.EncodeToString(crypted)) // 加密后的数据： Fe3xQJZVALMxovBw4qNGLA==

	// 解密
	origData, err := AesCbcPkcs7Decrypt(crypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后的数据：", string(origData)) // 解密后的数据： Hello, world!
}
