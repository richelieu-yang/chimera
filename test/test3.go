package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
)

func main() {
	data, err := base64Kit.EncodeFile("/Users/richelieu/Documents/ino/notes/Golang/WEB")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
