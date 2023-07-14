package sliceKit

import (
	"fmt"
	"testing"
)

func TestClear(t *testing.T) {
	s := make([]interface{}, 3, 6)
	s[0] = 0
	s[1] = 1
	s[2] = 2
	fmt.Println(s, len(s), cap(s))

	Clear(s)
	fmt.Println(s, len(s), cap(s))
}
