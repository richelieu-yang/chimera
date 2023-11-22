package caesarKit

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plainText := "测试 abcDEF"
	var shift int = 8

	cipherText := Encrypt(plainText, shift)
	decryptedText := Decrypt(cipherText, shift)
	fmt.Println("cipherText:", cipherText)
	fmt.Println("decryptedText:", decryptedText)

	if plainText != decryptedText {
		panic("not equal")
	} else {
		fmt.Println("equal")
	}
}

func TestEncryptAll(t *testing.T) {
	check := func(shift int) {
		plainText := "测试 abcDEF"
		cipherText := Encrypt(plainText, shift)
		decryptedText := Decrypt(cipherText, shift)

		if plainText != decryptedText {
			panic("not equal")
		}
	}

	for i := 0; i <= 100; i++ {
		fmt.Println("shift: ", i)
		check(i)
	}
}

//func TestEncryptWithRawURLBase64(t *testing.T) {
//	plainText := "测试 ~！@#￥%……&*（）——+·1234567890-=【】、「」|；‘。、，：“《》？abcDEF"
//	//plainText := "111000"
//	var shift uint8 = 8
//
//	cipherText := EncryptWithRawURLBase64(plainText, shift)
//	decryptedText, err := DecryptWithRawURLBase64(cipherText, shift)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("cipherText:", cipherText)
//	fmt.Println("decryptedText:", decryptedText)
//
//	if plainText != decryptedText {
//		panic("not equal")
//	} else {
//		fmt.Println("equal")
//	}
//}
