package main

import (
	"fmt"
)

func caesarCipherEncrypt(s string, shift int) string {
	shift = shift % 26
	runes := []rune(s)
	for i, r := range runes {
		switch {
		case 'a' <= r && r <= 'z':
			runes[i] = 'a' + (r-'a'+rune(shift)+26)%26
		case 'A' <= r && r <= 'Z':
			runes[i] = 'A' + (r-'A'+rune(shift)+26)%26
		}
	}
	return string(runes)
}

func caesarCipherDecrypt(s string, shift int) string {
	return caesarCipherEncrypt(s, -shift)
}

func main() {
	s := "Hello, World!"
	shift := -1
	encrypted := caesarCipherEncrypt(s, shift)
	fmt.Println("Encrypted:", encrypted)
	decrypted := caesarCipherDecrypt(encrypted, shift)
	fmt.Println("Decrypted:", decrypted)
}
