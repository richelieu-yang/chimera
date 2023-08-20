package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBuffer(nil)
	str := buffer.String() // ""
	fmt.Println(str)
}
