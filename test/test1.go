package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"path/filepath"
)

func main() {
	//path := ""
	//
	//fmt.Println(IsHidden(path))

	fmt.Println(fileKit.GetName(""))
	fmt.Println(fileKit.GetName(" "))  // " "
	fmt.Println(fileKit.GetName("  ")) // "  "
	fmt.Println(fileKit.GetName("."))
	fmt.Println(fileKit.GetName("./"))
	fmt.Println(fileKit.GetName("../"))

}

// IsHidden 文件（或目录）是否隐藏？
/*
e.g.
("") => false, nil
*/
func IsHidden(path string) (bool, error) {
	name := filepath.Base(path)
	return name[0:1] == ".", nil
}
