package main

import (
	"github.com/richelieu42/go-scales/src/assertKit"
)

func main() {
	test()
}

func test() {
	var obj interface{} = nil
	if err := assertKit.NotNil(obj, "obj"); err != nil {
		panic(err)
	}
}
