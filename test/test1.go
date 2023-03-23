package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var i int

	once.Do(func() {
		i++
		if true {
			return // 中断此匿名函数
		}
		i++
	})
	fmt.Println(i) // 1
}
