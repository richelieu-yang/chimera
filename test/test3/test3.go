package main

import (
	"fmt"
	"github.com/bytedance/sonic"
)

func main() {
	var str string = `{"a":"1","b":true}`

	{
		node, err := sonic.Get([]byte(str), "a")
		if err != nil {
			panic(err)
		}
		fmt.Println(node.String())
	}

	{
		node, err := sonic.Get([]byte(str), "b")
		if err != nil {
			panic(err)
		}
		fmt.Println(node.Bool())
	}
}
