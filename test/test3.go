package main

import (
	"fmt"
	"github.com/stevenyao/go-opencc"
)

const (
	input      = "中国鼠标软件打印机后台湾，参数错误: %s"
	config_s2t = "s2t.json"
)

func main() {
	fmt.Println("Test Converter class:")
	c := opencc.NewConverter(config_s2t)
	defer c.Close()
	output := c.Convert(input)
	fmt.Println(output)
}
