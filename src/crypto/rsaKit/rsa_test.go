package rsaKit

import (
	"fmt"
	"testing"
)

// 加密 && 解密
func TestRSA(t *testing.T) {
	options := []RsaOption{
		WithFormat(PKCS8),
		WithPassword("cyy"),
	}
	pri, pub, err := GenerateKeys(4096, options...)
	if err != nil {
		panic(err)
	}

	data := []byte("`1245`152678`169837~!@#$%^&*()_+=")
	// 加密
	data, err = Encrypt(data, pub)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文:", string(data))
	// 解密
	data, err = Decrypt(data, pri, options...)
	if err != nil {
		panic(err)
	}
	fmt.Println("明文:", string(data))
}
