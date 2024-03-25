package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io/fs"
	"path/filepath"
	"regexp"
)

func main() {
	re := regexp.MustCompile("\\[(\\d+)[vV]*(\\d*)\\]")

	err := filepath.WalkDir("/Users/richelieu/Downloads/魔都精兵的奴隶", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !fileKit.IsFile(path) {
			return nil
		}

		name := fileKit.GetName(path)
		ext := fileKit.GetExt(path)
		s := re.FindAllStringSubmatch(name, -1)
		newName := s[0][1] + ext

		//fileKit.Rename()

		fmt.Println(newName)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
