package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
	"time"
)

func main() {
	set := setKit.NewSet(false, 0, 1, 2, 3)
	set.Each(func(i int) bool {
		go func() {
			fmt.Println("a", i)
		}()
		return false
	})
	time.Sleep(time.Second)

	s := []int{0, 1, 2, 3}
	for _, ele := range s {
		go func() {
			fmt.Println("b", ele)
		}()
	}
	time.Sleep(time.Second)
}
