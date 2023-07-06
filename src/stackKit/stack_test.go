package stackKit

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int](true)
	for i := 0; i < 3; i++ {
		var i2 = i
		stack.Push(i2)
	}
	for i := 0; i < 4; i++ {
		fmt.Println(stack.Pop())
	}
}
