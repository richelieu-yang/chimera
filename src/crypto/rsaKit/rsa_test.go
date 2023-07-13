package rsaKit

import (
	"fmt"
	"testing"
)

// 加密 && 解密
func TestRSA(t *testing.T) {
	password := "qwdqwdqwd"
	options := []RsaOption{
		WithFormat(PKCS1),
		WithPassword(password),
	}
	pri, pub, err := GenerateKeys(2048, options...)
	if err != nil {
		panic(err)
	}

	data := []byte("`1245`152678`169837~!@#$%^&*()_+=")
	// 加密
	data, err = Encrypt(data, pub)
	if err != nil {
		panic(err)
	}
	fmt.Printf("密文:\n%s\n", string(data))
	// 解密
	data, err = Decrypt(data, pri, password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("明文:\n%s\n", string(data))
}
