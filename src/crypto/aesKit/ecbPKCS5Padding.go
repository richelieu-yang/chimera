package aesKit

import "github.com/zeromicro/go-zero/core/codec"

// EncryptEcbPKCS5PaddingToString
/*

 */
var EncryptEcbPKCS5PaddingToString func(key, plainText string) (string, error) = codec.EcbEncryptBase64

// DecryptEcbPKCS5PaddingFromString
/*

 */
var DecryptEcbPKCS5PaddingFromString func(key, cipherText string) (string, error) = codec.EcbDecryptBase64

// EncryptEcbPKCS5Padding
/*

 */
var EncryptEcbPKCS5Padding func(key, plainText []byte) ([]byte, error) = codec.EcbEncrypt

// DecryptEcbPKCS5Padding
/*

 */
var DecryptEcbPKCS5Padding func(key, cipherData []byte) ([]byte, error) = codec.EcbDecrypt
