package unicodeKit

import (
	"fmt"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	str := "软心姑娘sss"

	encoded := Encode(str)
	decoded, err := Decode(encoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(encoded)
	fmt.Println(decoded)

	if str != decoded {
		panic("not equal")
	}
}
