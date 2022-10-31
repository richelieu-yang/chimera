package main

import (
	"github.com/richelieu42/go-scales/src/core/ioKit"
)

func main() {
	wc, err := ioKit.NewRotateFileWriteCloser1("ccc.log", -1, -1, true)
	if err != nil {
		panic(err)
	}

	if _, err := wc.Write([]byte("abc")); err != nil {
		panic(err)
	}
	if err := wc.Close(); err != nil {
		panic(err)
	}
	if _, err := wc.Write([]byte("def")); err != nil {
		panic(err)
	}
}
