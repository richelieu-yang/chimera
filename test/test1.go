package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"os"
)

func main() {
	fileKit.NewFile()

	dirPath := ""
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------------------------")
}
