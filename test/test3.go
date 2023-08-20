package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBuffer(nil)
	buffer.Write
	str := buffer.String()
	fmt.Println(str)
}
