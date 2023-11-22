package morseKit

import (
	"github.com/alwindoss/morse"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
)

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
