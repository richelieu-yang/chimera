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
		fmt.Println(str, len(str))
	}
}
