package main

import (
	"fmt"
	"os"
)

func main() {
	dirPath := ""
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------------------------")
}
