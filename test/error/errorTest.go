package main

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"os"
)

// 多次调用errorKit.WithStackOnce()的话，只有第一次有效
func main() {
	_, err := os.ReadFile("")
	if err != nil {
		err = errorKit.WithStackOnce(err)
		err = errorKit.WithStackOnce(err)
		err = errorKit.WithStackOnce(err)
		fmt.Printf("%+v", err)
	}
}
