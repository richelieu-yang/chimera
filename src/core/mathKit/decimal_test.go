package mathKit

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	fmt.Println(Round(3.14, 1))  // 3.1
	fmt.Println(Round(3.15, 1))  // 3.2
	fmt.Println(Round(-3.14, 1)) // -3.1
	fmt.Println(Round(-3.15, 1)) // -3.2
}
