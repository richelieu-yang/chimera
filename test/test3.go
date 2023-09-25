package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
)

func main() {
	url, err := urlKit.Parse("http://localhost:8848/nacos")
	if err != nil {
		panic(err)
	}
	fmt.Println(url.Hostname())
	fmt.Println(url.Port())
	fmt.Println(url.Scheme)
	fmt.Println(url.Path)
}
