package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	println(fileKit.GetExt("main.go"))  // ".go"
	println(fileKit.GetExt("api.json")) // ".json"
	println(fileKit.GetExt(""))         // ""
	println(fileKit.GetExt("    "))     // ""
	println(fileKit.GetExt("empty"))    // ""
}
