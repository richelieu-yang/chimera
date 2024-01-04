// Package aesKit
/*
更多可以参考：https://github.com/hwholiday/learning_tools/blob/master/encryption_algorithm/aes.go
*/
package aesKit1

import "bytes"

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, paddingText...)
}

func ZeroUnpadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	paddingText := int(origData[length-1])
	return origData[:(length - paddingText)]
}
