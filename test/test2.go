package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
)

func main() {
	if err := fileKit.MkDirs("."); err != nil {
		panic(err)
	}
	if err := fileKit.MkDirs("./"); err != nil {
		panic(err)
	}
	if err := fileKit.MkDirs("/"); err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------------")
}
