package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
)

func main() {
	writeCloser, err := ioKit.NewRotatableWriteCloser("aaa.log", dataSizeKit.EiB)
	if err != nil {
		panic(err)
	}

	writeCloser.Rotate()
	writeCloser.Rotate()
	writeCloser.Rotate()
}
