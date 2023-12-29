package dataSizeKit

import (
	"fmt"
	"testing"
)

func TestByteToMiB(t *testing.T) {
	bytes := MiB*10 + KiB*123
	fmt.Println(ByteToMiB(bytes, 5)) // 10.12012
}
