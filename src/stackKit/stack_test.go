package stackKit

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int](false)
	for i := 0; i < 3; i++ {
		stack.Push(i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(stack.Pop())
	}
}
