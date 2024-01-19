package brotliKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/bytesKit"
	"testing"
)

func TestEntire(t *testing.T) {
	data := []byte("Don't communicate by sharing memory, share memory by communicating.Don't communicate by sharing memory, share memory by communicating.Don't communicate by sharing memory, share memory by communicating.")
	fmt.Println(len(data))

	// Compress
	compressed, err := Compress(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("compressed:", string(compressed))
	fmt.Println(len(compressed))

	// Decompress
	decompressed, err := Decompress(compressed)
	if err != nil {
		panic(err)
	}
	fmt.Println("decompressed:", string(decompressed))
	fmt.Println(len(decompressed))

	if !bytesKit.Equals(data, decompressed) {
		panic("not equals")
	}
}
