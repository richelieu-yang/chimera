package main

import (
	"fmt"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
)

func main() {
	padding := 0
	data, err := gaes.EncryptCFB([]byte("111"), []byte("1234567891234567"), &padding)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64Kit.EncodeToString(data))
}
