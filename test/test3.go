package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"io/fs"
)

func main() {
	dirPath := "/Users/richelieu/Desktop/a"

	err := pathKit.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("======")

	err = pathKit.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
