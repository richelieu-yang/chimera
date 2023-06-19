package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"os"
)

func main() {
	f, err := fileKit.NewFile("aaaa.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	os.Stdout = f
	fmt.Println("aaa\n")
}
