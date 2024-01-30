package mathKit

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println(Max(0, -1, 2, 300, 4)) // 300
}

func TestMin(t *testing.T) {
	fmt.Println(Min(0, -1, 2, 300, 4)) // -1
}
