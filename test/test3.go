package main

import (
	"github.com/natefinch/lumberjack/v3"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	wc, err := lumberjack.NewRoller("ccc.log", 20*dataSizeKit.MiB, &lumberjack.Options{})
	if err != nil {
		panic(err)
	}

	//wc := &lumberjack.Logger{
	//	Filename: "CCC.log",
	//	MaxSize:  20,
	//}
	for {
		n, err := wc.Write([]byte("46464\n"))
		if err != nil {
			logrus.Fatal(err)
			return
		}
		logrus.Info(n)
	}
}
