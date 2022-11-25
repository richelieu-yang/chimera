package main

import (
	"github.com/richelieu42/go-scales/src/core/ioKit"
	"io"
	"time"
)

func main() {
	wc, err := ioKit.NewRotateFileWriteCloser("aaa.log", time.Second*3, time.Second*30, true)
	if err != nil {
		panic(err)
	}

	for {
		_, err := io.WriteString(wc, time.Now().String()+"\n")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
