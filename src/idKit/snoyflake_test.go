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
	for i := 0; i < 100; i++ {
		fmt.Println(sf.NextID())
	}
}
