package main

import (
	"github.com/richelieu42/go-scales/src/core/ioKit"
)

func main() {
	wc, err := ioKit.NewFileWriterCloser("a.log", true)
	if err != nil {
		panic(err)
	}

	_, err = wc.Write([]byte("123"))
	if err != nil {
		panic(err)
	}
	if err := wc.Close(); err != nil {
		panic(err)
	}
	_, err = wc.Write([]byte("456"))
	if err != nil {
		panic(err)
	}
}
