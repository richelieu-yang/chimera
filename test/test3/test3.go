package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/versionKit"
)

func main() {
	v, err := versionKit.NewVersion("1.3.10+meta")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
	fmt.Println(v.String())

	//var str string = `{"a":"1","b":true}`
	//
	//{
	//	node, err := sonic.Get([]byte(str), "a")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(node.String())
	//}
	//
	//{
	//	node, err := sonic.Get([]byte(str), "b")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(node.Bool())
	//}
}
