package caesarKit

import "bytes"

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
