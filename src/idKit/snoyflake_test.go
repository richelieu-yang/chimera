package idKit

import (
	"fmt"
	"testing"
)

func TestNewSonyFlake(t *testing.T) {
	sf, err := NewSonyFlake(nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i >= 0; i++ {
		id, err := sf.NextID()
		if err != nil {
			panic(err)
		}

		str := fmt.Sprintf("%d", id)
		if len(str) != 18 {
			panic("len(str) != 18")
		}
		fmt.Println(str)
	}
}
