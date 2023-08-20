package caesarKit

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
