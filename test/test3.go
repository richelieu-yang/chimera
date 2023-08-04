package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	dirPath := "/Users/richelieu/Desktop/a"

	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
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

	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
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
