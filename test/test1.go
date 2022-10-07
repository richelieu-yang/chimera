package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/http/httpClientKit"
)

func main() {
	data, err := httpClientKit.Post("https://cn.bing.com/1/1", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println("================================================")
}
