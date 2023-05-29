package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	options := []rsaKit.RsaOption{
		rsaKit.WithPassword(""),
		rsaKit.WithFormat(rsaKit.PKCS1),
	}

	priPath := "pri.pem"
	pubPath := "pub.pem"
	if err := rsaKit.GenerateKeyFiles(4096, priPath, pubPath, options...); err != nil {
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
	data, err = rsaKit.Encrypt(data, pub)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文:", string(data))
	// 解密
	data, err = rsaKit.Decrypt(data, pri, options...)
	if err != nil {
		panic(err)
	}
	fmt.Println("明文:", string(data))
}
