package main

import (
	"encoding/base64"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
)

func main() {
	str, err := base64Kit.DecodeStringToString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", base64Kit.WithEncoding(base64.RawURLEncoding))
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
