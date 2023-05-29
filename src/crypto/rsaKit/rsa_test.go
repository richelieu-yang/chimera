package rsaKit

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"testing"
)

func TestRSA(t *testing.T) {
	options := []RsaOption{
		WithPassword("qwdqwdqwdqwdqwdqwdqwdq"),
		WithFormat(PKCS8),
	}

	priPath := "_pri.key"
	pubPath := "_pub.key"
	if err := GenerateKeyFiles(4096, priPath, pubPath, options...); err != nil {
		panic(err)
	}

	pri, err := fileKit.ReadFile(priPath)
	if err != nil {
		panic(err)
	}
	pub, err := fileKit.ReadFile(pubPath)
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
