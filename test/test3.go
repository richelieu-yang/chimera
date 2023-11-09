package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
)

func main() {
	m := mapKit.NewMapWithLock[string, interface{}]()
	m.LockFunc(func() {
		fmt.Println(len(m.Map))
	})
}
