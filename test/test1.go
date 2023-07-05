package main

import (
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

func main() {
	sliceKit.ConvertElementType()

	list := glist.New()
	list.PushBack(0)
	list.PushBack(1)

	fmt.Println(list.PopBack())
	fmt.Println(list.PopBack())
	fmt.Println(list.PopBack())
}
