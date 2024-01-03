package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// AES-CBC-PKCS7 加解密，支持传参 iv
// key 必须是 16 (AES-128)、24 (AES-192) 或 32 (AES-256) 字节的 AES 密钥
// iv 必须是 16 字节的初始化向量

// PKCS7Padding 填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding 去除填充
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// AesEncrypt AES-CBC-PKCS7 加密
func AesEncrypt(data []byte, key []byte, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	data = PKCS7Padding(data, blockSize)
	crypted := make([]byte, len(data))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted, data)
	return crypted
}

// AesDecrypt AES-CBC-PKCS7 解密
func AesDecrypt(data []byte, key []byte, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	origData = PKCS7UnPadding(origData)
	return origData
}

func main() {
	// 测试数据
	data := []byte("Hello, world!")
	key := []byte("00000yozo_config")
	iv := []byte("00000yozo_config")

	// 加密
	crypted := AesEncrypt(data, key, iv)
	fmt.Println("加密后的数据：", base64.StdEncoding.EncodeToString(crypted)) // 加密后的数据： Fe3xQJZVALMxovBw4qNGLA==

	// 解密
	origData := AesDecrypt(crypted, key, iv)
	fmt.Println("解密后的数据：", string(origData)) // 解密后的数据： Hello, world!
}
