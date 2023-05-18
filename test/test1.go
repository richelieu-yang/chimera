package main

import (
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/richelieu42/chimera/v2/src/jsonKit"
)

func main() {
	l := glist.NewFrom([]interface{}{0, 1, 2})

	// 序列化
	json, err := jsonKit.MarshalToString(l)
	if err != nil {
		panic(err)
	}
	fmt.Println(json) // [0,1,2]

	// 反序列化
	l1 := &glist.List{}
	err = jsonKit.UnmarshalFromString(json, l1)
	if err != nil {
		panic(err)
	}
	fmt.Println(l1) // [0,1,2]
}
