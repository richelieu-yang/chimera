package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/ocr/gosseractKit"
)

func main() {
	text, err := gosseractKit.GertText("")
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}
