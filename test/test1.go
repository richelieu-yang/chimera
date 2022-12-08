package main

import (
	"github.com/richelieu42/go-scales/src/component/componentKit"
)

func main() {
	err := componentKit.InitializeEnvironment()
	if err != nil {
		panic(err)
	}
}
