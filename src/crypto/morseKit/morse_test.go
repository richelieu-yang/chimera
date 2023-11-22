package morseKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	text := `测试`
	text = string(base64Kit.Encode([]byte(text)))
	text = text + "+-/_="
	fmt.Println(text)

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

	if !strKit.EqualsIgnoreCase(string(plainData), text) {
		panic("not equal")
	}

	//if string(plainData) != text {
	//	panic("not equal")
	//}

	fmt.Println(base64Kit.Decode([]byte("5RWL6K+V+-/_=")))

}
