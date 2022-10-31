package main

import (
	"github.com/richelieu42/go-scales/src/core/ioKit"
	"github.com/sirupsen/logrus"
)

func main() {
	wc, err := ioKit.NewRotateFileWriteCloser1("ccc.log", -1, -1, true)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.Out = wc

	logger.Warn("16346498498")
	if err := wc.Close(); err != nil {
		panic(err)
	}
	logger.Warn("强无敌群无多群无多")

	//if _, err := wc.Write([]byte("abc")); err != nil {
	//	panic(err)
	//}
	//if err := wc.Close(); err != nil {
	//	panic(err)
	//}
	//if _, err := wc.Write([]byte("def")); err != nil {
	//	panic(err)
	//}
}
