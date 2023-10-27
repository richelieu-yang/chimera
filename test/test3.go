package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

func main() {
	m := mutexKit.NewRWMutex()

	m.LockFunc(func() {
		fmt.Println(666)
	})
}
