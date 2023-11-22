package unicodeKit

import "testing"

func TestEncodeAndDecode(t *testing.T) {
	str := "软心姑娘"

	encoded := Encode(str)
	decoded, err := Decode(encoded)
	if err != nil {
		panic(err)
	}
	if str != decoded {
		panic("not equal")
	}
}
