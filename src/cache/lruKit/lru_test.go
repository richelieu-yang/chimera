package lruKit

import (
	"fmt"
	"testing"
)

func TestNewCache(t *testing.T) {
	l, _ := NewCache[int, any](8)
	for i := 0; i < 16; i++ {
		l.Add(i, nil)
	}
	fmt.Println(l.Len()) // 8
}
