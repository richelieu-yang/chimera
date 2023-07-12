package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbase64"
)

func main() {

	gbase64.Decode()

	input := []byte("\x00\x00\x3e\x00\x00\x3f\x00")
	fmt.Println(string(input))
	fmt.Println(base64.StdEncoding.EncodeToString(input))                               // AAA+AAA/AA==
	fmt.Println(base64.URLEncoding.EncodeToString(input))                               // AAA-AAA_AA==
	fmt.Println(base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(input)) // AAA+AAA/AA
	fmt.Println(base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(input)) // AAA-AAA_AA
}
