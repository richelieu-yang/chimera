package atomicKit

import (
	"fmt"
	"testing"
)

func TestNewInt32(t *testing.T) {
	i := NewInt32(0)

	fmt.Println(i.Add(1)) // 1
	fmt.Println(i.Add(1)) // 2
	fmt.Println(i.Load()) // 2
	fmt.Println("======")

	/* CompareAndSwap */
	fmt.Println(i.CompareAndSwap(2, 3)) // true
	fmt.Println(i.Load())               // 3
	fmt.Println(i.CompareAndSwap(2, 3)) // false
	fmt.Println(i.Load())               // 3
	fmt.Println("======")

	/* Store */
	i.Store(666)
	fmt.Println(i.Load()) // 666
}
