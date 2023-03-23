package main

import (
	"github.com/richelieu42/chimera/src/core/ioKit"
	"github.com/sirupsen/logrus"
)

func main() {
	writeCloser, err := ioKit.NewLumberjackWriteCloser("/Users/richelieu/Downloads/a.log", 1, 10, 7, true, false)
	if err != nil {
		logrus.Fatal(err)
	}

	//for {
	//	_, _ = writeCloser.Write([]byte(idKit.NewULID() + "\r\n"))
	//	time.Sleep(time.Second)
	//}

	_, _ = writeCloser.Write([]byte("0\n"))
	_, _ = writeCloser.Write([]byte("1\n"))
	_, _ = writeCloser.Write([]byte("2\n"))
	_, _ = writeCloser.Write([]byte("3\n"))
}
