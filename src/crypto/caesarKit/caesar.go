package caesarKit

import (
	"bytes"
	"encoding/base64"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
)

// Encrypt 凯撒密码加密.
/*
PS: 仅加密大小写的英文字母.

@param shift 推荐值: (1) [0, 25]
					(2) 不推荐使用 26*n(n >= 0)，因为这样加密了个寂寞
*/
func Encrypt(plainText string, shift uint8) string {
	shift = polyfillShift(shift)

	buffer := bytes.NewBuffer(nil)
	for _, char := range plainText {
		if char >= 'a' && char <= 'z' {
			char = (char-'a'+rune(shift))%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			char = (char-'A'+rune(shift))%26 + 'A'
		}
		buffer.WriteRune(char)
	}
	return buffer.String()
}

func Decrypt(cipherText string, shift uint8) string {
	shift = polyfillShift(shift)

	return Encrypt(cipherText, 26-shift)
}

func polyfillShift(shift uint8) uint8 {
	// 确保偏移量在0-25之间
	return shift % 26
}

func EncryptWithBase64(str string, shift uint8) string {
	plainText := base64Kit.EncodeStringToString(str, base64Kit.WithEncoding(base64.RawURLEncoding))

	return Encrypt(plainText, shift)
}

func DecryptWithBase64(str string, shift uint8) (string, error) {
	cipherText := Decrypt(str, shift)

	decryptedText, err := base64Kit.DecodeStringToString(cipherText, base64Kit.WithEncoding(base64.RawURLEncoding))
	if err != nil {
		return "", err
	}
	return decryptedText, nil
}
