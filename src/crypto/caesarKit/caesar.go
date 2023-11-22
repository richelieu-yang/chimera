package caesarKit

import "github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"

// Encrypt 凯撒密码加密.
/*
PS: 仅加密大小写的英文字母.

@param shift 推荐值: (1) [0, 25]
					(2) 不推荐使用 26*n(n >= 0)，因为这样加密了个寂寞
*/
func Encrypt(plainText string, shift int) (cipherText string) {
	cipher := NewCaesarCipher(shift)
	cipherText = cipher.Encrypt(plainText)
	return
}

func Decrypt(cipherText string, shift int) (plainText string) {
	cipher := NewCaesarCipher(shift)
	plainText = cipher.Decrypt(cipherText)
	return
}

func EncryptWithBase64(plainText string, shift int) (cipherText string) {
	// base64 编码
	base64Str := base64Kit.EncodeStringToString(plainText)
	// 凯撒密码 加密
	cipherText = Encrypt(base64Str, shift)
	return
}

func DecryptWithBase64(cipherText string, shift int) (plainText string, err error) {
	// 凯撒密码 解密
	base64Str := Decrypt(cipherText, shift)
	// base64 解码
	plainText, err = base64Kit.DecodeStringToString(base64Str)
	return
}
