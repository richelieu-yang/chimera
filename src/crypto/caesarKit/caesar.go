package caesarKit

// Encrypt 凯撒密码加密.
/*
PS:
(1) 仅加密大小写的英文字母;
(2) 可以搭配 base64 使用.

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
