package main

import (
	"encoding/base64"
	"fmt"
	"unsafe"
)

func main() {
	enc := base64.StdEncoding
	fmt.Println(unsafe.Pointer(enc))
	enc1 := enc.WithPadding(base64.NoPadding)
	fmt.Println(unsafe.Pointer(enc1))

}
