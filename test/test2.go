package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := io.LimitedReader{
		R: strings.NewReader("0123456789abcdef"),
		N: 6,
	}
	s := make([]byte, 3)
	fmt.Println(reader.Read(s))
	fmt.Println(string(s))
	fmt.Println(reader.Read(s))
	fmt.Println(string(s))
	fmt.Println(reader.Read(s))
	fmt.Println(string(s))
}
