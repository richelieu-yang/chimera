package aesKit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AesCbcPkcs7Encrypt AES/CBC/PKCS7 加密.
/*
PS: 返回值转换为字符串，可以考虑使用: base64.StdEncoding.EncodeToString().
*/
func AesCbcPkcs7Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)
	crypted := make([]byte, len(data))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

//func AesCbcPkcs7EncryptToBase64(data []byte, key []byte, iv []byte) ([]byte, error) {
//	data, err := AesCbcPkcs7Encrypt(data, key, iv)
//	if err != nil {
//		return "", err
//	}
//	return base64.StdEncoding.Encode(data), nil
//}
//
//func AesCbcPkcs7EncryptToHex(data []byte, key []byte, iv []byte) ([]byte, error) {
//	data, err := AesCbcPkcs7Encrypt(data, key, iv)
//	if err != nil {
//		return "", err
//	}
//	return hex.Encode(data), nil
//}

// AesCbcPkcs7Decrypt AES/CBC/PKCS7 解密.
func AesCbcPkcs7Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	origData = pkcs7UnPadding(origData)
	return origData, nil
}

//func AesCbcPkcs7DecryptFromBase64(data []byte, key []byte, iv []byte) ([]byte, error) {
//
//}
//
//func AesCbcPkcs7DecryptFromHex(data []byte, key []byte, iv []byte) ([]byte, error) {
//
//}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 去除填充
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}
