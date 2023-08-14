package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	path := "/Users/richelieu/Downloads"
	entries, err := fileKit.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
