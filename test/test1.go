package main

import (
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/sirupsen/logrus"
)

func main() {
	//tmp := uint64(dataSizeKit.MB.GetValue() * 87)
	//
	//size := humanize.ToReadableStringWithSI(tmp)
	//fmt.Printf("That file is %s.\n", size)

	//data, err := fileKit.ReadFile("/Users/richelieu/Documents/ino/notes/Golang/gorm.wps")
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//length := len(data)
	//
	////length := dataSizeKit.KB.GetValue() * 79
	//
	//logrus.Info(length)
	//logrus.Info(humanize.Bytes(uint64(length)))

	logrus.Info(dataSizeKit.ToReadableStringWithSI(78848))
	logrus.Info(dataSizeKit.ToReadableStringWithIEC(78848))
}
