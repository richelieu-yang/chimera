package main

import (
	"fmt"
)

func Encrypt(plaintext string, key int) string {
	ciphertext := ""

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			char = (char-'a'+rune(key))%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			char = (char-'A'+rune(key))%26 + 'A'
		}
		ciphertext += string(char)
	}
	return ciphertext
}

func Decrypt(ciphertext string, key int) string {
	return Encrypt(ciphertext, 26-key)
}

//func encryptCaesar(plaintext string, key int) string {
//	ciphertext := ""
//
//	for _, char := range plaintext {
//		if char >= 'a' && char <= 'z' {
//			char = (char-'a'+rune(key))%26 + 'a'
//		} else if char >= 'A' && char <= 'Z' {
//			char = (char-'A'+rune(key))%26 + 'A'
//		}
//		ciphertext += string(char)
//	}
//	return ciphertext
//}
//
//func decryptCaesar(ciphertext string, key int) string {
//	return encryptCaesar(ciphertext, 26-key)
//}

func main() {
	plaintext := "hello world"
	key := 3

	ciphertext := Encrypt(plaintext, key)
	fmt.Println("cipher text:", ciphertext)
	decryptedText := Decrypt(ciphertext, key)
	fmt.Println("decrypted text:", decryptedText)
}
