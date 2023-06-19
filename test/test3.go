package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"time"
)

func main() {
	path := "test.log"

	f, err := fileKit.NewFileInAppendMode(path)
	if err != nil {
		panic(err)
	}
	writeCloser, err := ioKit.NewDailyWriteCloser(path)
	if err != nil {
		panic(err)
	}

	cf, err := fileKit.NewCustomizedFile(f, writeCloser)
	if err != nil {
		panic(err)
	}

	for {
		_, err := cf.WriteString("qwdqwdqwdqwdqwdqwdqwdqwdqw\r\n")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
