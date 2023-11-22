package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/crypto/morseKit"
)

func main() {
	text := "test测试"
	//text = unicodeKit.Encode(text)
	//fmt.Println(text)

	morsedData, err := morseKit.Encode([]byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(morsedData))
}
