package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrusKit.MustSetUp(nil)

	wc, err := ioKit.NewRotatableWriteCloser("/Users/richelieu/Downloads/aaa.log", 1024*1024*1,
		ioKit.WithCompress(true),
		ioKit.WithMaxAge(time.Minute),
	)
	if err != nil {
		panic(err)
	}

	for {
		n, err := wc.Write([]byte("qwdqwdqwdqwd\n"))
		if err != nil {
			logrus.Fatal(err)
			return
		}
		logrus.Info(n)
	}
}
