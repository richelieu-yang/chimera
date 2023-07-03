package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	println(fileKit.GetExtName("main.go"))  // "go"
	println(fileKit.GetExtName("api.json")) // "json"
	println(fileKit.GetExtName(""))         // ""
	println(fileKit.GetExtName("    "))     // ""
	println(fileKit.GetExtName("empty"))    // ""
}
