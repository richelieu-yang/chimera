package caesarKit

import (
	"fmt"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	plainText := "测试 abcDEF"
	var shift uint8 = 100

	cipherText := Encrypt(plainText, shift)
	decryptedText := Decrypt(cipherText, shift)

	fmt.Println("cipherText: ", cipherText)
	fmt.Println("decryptedText: ", decryptedText)
	if plainText != decryptedText {
		panic("not equal")
	} else {
		fmt.Println("equal")
	}
}
