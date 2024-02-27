package main

import (
	"fmt"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println(language.Make("zh"))
	fmt.Println(language.Make("zh-CN"))
	fmt.Println(language.Make("zh-HK"))
}
