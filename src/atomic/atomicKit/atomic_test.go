package atomicKit

import (
	"fmt"
	"testing"
)

func TestNewInt32(t *testing.T) {
	i := NewInt32(0)

	i.Add(1)
	i.Add(1)
	fmt.Println(i.Load()) // 2
	fmt.Println("======")

	fmt.Println(i.CompareAndSwap(2, 3)) // true
	fmt.Println(i.Load())               // 3

	fmt.Println(i.CompareAndSwap(2, 3)) // false
	fmt.Println(i.Load())               // 3
}
