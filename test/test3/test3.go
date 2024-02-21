package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/versionKit"
)

func main() {
	v, err := versionKit.NewVersion("1.0.0-alpha+001")
	if err != nil {
		panic(err)
	}
	/*
		v.pre == "alpha"
		v.metadata == "001"
	*/
	fmt.Println(v)
}
