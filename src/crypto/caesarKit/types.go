package caesarKit

type CaesarCipher struct {
	Shift int
}

func (cipher *CaesarCipher) Encrypt(plainText string) (cipherText string) {
	cipherText = common(plainText, cipher.Shift)
	return
}

func (cipher *CaesarCipher) Decrypt(cipherText string) (plainText string) {
	plainText = common(cipherText, -cipher.Shift)
	return
}

func common(str string, shift int) string {
	// 确保偏移量在0-25之间
	shift = shift % 26

	runes := []rune(str)
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

// NewCaesarCipher
/*
@param Shift 	(1) 推荐值: [1, 25]
				(2) 不推荐使用 26*n(n >= 0)，因为这样加密了个寂寞
				(3) 可以是负数，表示向左偏移
*/
func NewCaesarCipher(shift int) *CaesarCipher {
	// 确保偏移量在0-25之间
	shift = shift % 26

	return &CaesarCipher{
		Shift: shift,
	}
}
