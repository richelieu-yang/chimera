package morseKit

import (
	"fmt"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	text := "Hello, world"

	plainData := []byte(text)
	cryptoData, err := Encode(plainData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("cryptoText: ", string(cryptoData))

	plainData, err = Decode(cryptoData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("plainText: ", string(plainData))

	if string(plainData) != text {
		panic("not equal")
	}
}
