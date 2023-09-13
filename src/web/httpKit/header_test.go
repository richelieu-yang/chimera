package httpKit

import (
	"fmt"
	"testing"
)

func TestAddHeader(t *testing.T) {
	header := make(map[string][]string)

	AddHeader(header, "k", "0")
	fmt.Println(header) // map[K:[0]]
	AddHeader(header, "k", "1")
	fmt.Println(header) // map[K:[0 1]]
	AddHeader(header, "k", "1")
	fmt.Println(header) // map[K:[0 1 1]]
}
