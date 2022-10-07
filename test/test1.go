package main

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/http/httpClientKit"
)

func main() {
	data, err := httpClientKit.Post("https://cn.bing.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println("-----------------------------")
}
