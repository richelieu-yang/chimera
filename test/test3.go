package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"time"
)

func main() {
	wc, err := ioKit.NewDailyWriteCloser("/Users/richelieu/Downloads/a.log")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(wc.Write([]byte("qwdqwdqwdqwd\n")))
		time.Sleep(time.Second)
	}
}
