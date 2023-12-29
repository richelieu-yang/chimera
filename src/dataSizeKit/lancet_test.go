package dataSizeKit

import (
	"fmt"
	"testing"
)

func TestToReadableIecString(t *testing.T) {
	result1 := ToReadableIecString(1024)
	result2 := ToReadableIecString(1024 * 1024)
	result3 := ToReadableIecString(1234567)
	result4 := ToReadableIecString(1234567, 2)

	fmt.Println(result1) // 1KiB
	fmt.Println(result2) // 1MiB
	fmt.Println(result3) // 1.1774MiB
	fmt.Println(result4) // 1.18MiB
}
