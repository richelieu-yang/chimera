package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"time"
)

func main() {
	//path := "nohup.out"

	start := time.Now()

	path := "/Users/richelieu/GolandProjects/chimera/nohup111_副本5.out"

	if _, err := fileKit.Create(path); err != nil {
		panic(err)
	}

	if err := fileKit.Truncate(path, 0); err != nil {
		panic(err)
	}

	//if err := fileKit.CopyFile(path, "nohup1.out"); err != nil {
	//	panic(err)
	//}
	fmt.Println(time.Since(start))
	//if _, err := fileKit.Create(path); err != nil {
	//	panic(err)
	//}
}
