package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	pri, err := fileKit.ReadFile("pri.pem")
	if err != nil {
		panic(err)
	}
	pub, err := fileKit.ReadFile("pub.pem")
	if err != nil {
		panic(err)
	}

	data, err := rsaKit.PKCS1Encrypt([]byte("c强无敌群无多yy强无敌群"), pub)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data, err = rsaKit.PKCS1Decrypt(data, pri, []byte("cyy"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
