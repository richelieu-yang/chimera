package morseKit

import (
	"github.com/alwindoss/morse"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
)

// Encode
/*
@param plainData 支持中文
*/
func Encode(plainData []byte) (cryptoData []byte, err error) {
	reader := ioKit.NewReader(plainData)
	hacker := morse.NewHacker()
	return hacker.Encode(reader)
}

func Decode(cryptoData []byte) (plainData []byte, err error) {
	reader := ioKit.NewReader(cryptoData)
	hacker := morse.NewHacker()
	return hacker.Decode(reader)
}
