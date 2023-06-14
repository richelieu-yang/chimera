package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	wc, err := ioKit.NewDailyWriteCloser("aaa.log")
	if err != nil {
		logrus.Fatal(err)
	}
	logger := logrusKit.NewLogger(logrusKit.WithOutput(wc))
	c, _, err := cronKit.NewCronWithTask("* * * * * *", func() {
		logger.Info("-")
	})
	if err != nil {
		logrus.Fatal(err)
	}

	c.Run()

	//wc, err := ioKit.NewRotatableWriteCloser("/Users/richelieu/Downloads/aaa.log", 1024*1024*1,
	//	ioKit.WithCompress(true),
	//	ioKit.WithMaxAge(time.Minute),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for {
	//	n, err := wc.Write([]byte("qwdqwdqwdqwd\n"))
	//	if err != nil {
	//		logrus.Fatal(err)
	//		return
	//	}
	//	logrus.Info(n)
	//}
}
