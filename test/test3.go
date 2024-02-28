package main

import (
	"fmt"
	"github.com/stevenyao/go-opencc"
)

func main() {
	path := "./s2t.json"

	converter := opencc.NewConverter(path)
	defer converter.Close()

	output := converter.Convert("中国鼠标软件打印机后台湾，参数错误: %s")
	fmt.Println("繁体: ", output)
}
