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

func TestParseIecString(t *testing.T) {
	result1, _ := ParseIecString("12")
	result2, _ := ParseIecString("12ki")
	result3, _ := ParseIecString("12 KiB")
	result4, _ := ParseIecString("12.2 kib")

	fmt.Println(result1) // 12
	fmt.Println(result2) // 12288
	fmt.Println(result3) // 12288
	fmt.Println(result4) // 12492
}
