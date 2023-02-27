package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func remove(slice []int, elem int) []int {
	index := -1
	for i, e := range slice {
		if e == elem {
			index = i
			break
		}
	}
	if index == -1 {
		// 如果元素不存在于切片中，返回原始切片
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 移除元素
	fmt.Println(sliceKit.Remove(slice, 0))
}
