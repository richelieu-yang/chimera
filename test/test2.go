package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/http/httpKit"
)

type MyInt int

type Bean struct {
}

func main() {
	header := make(map[string][]string)

	fmt.Println(header) // map[]
	httpKit.AddHeader(header, "k", "0")
	fmt.Println(header) // map[K:[0]]
	httpKit.AddHeader(header, "k", "1")
	fmt.Println(header) // map[K:[0 1]]

	httpKit.SetHeader(header, "k", "2")
	fmt.Println(header) // map[K:[2]]
	httpKit.SetHeader(header, "k", "")
	fmt.Println(header) // map[]
}
