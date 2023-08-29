package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
)

func init() {
	cpuKit.SetUp()
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: true,
	})
}

func main() {
	//// 看下能否恢复
	//defer func() {
	//	if obj := recover(); obj != nil {
	//		logrus.Infof("recover: [%T, %v].", obj, obj)
	//	}
	//}()
	//
	//m := map[string]interface{}{
	//	"0": 3.1415926,
	//	"1": 1,
	//}
	//jsonStr, err := sonic.ConfigStd.MarshalToString(m)
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//logrus.Infof("jsonStr: [%s].", jsonStr)
	//var m1 map[string]interface{}
	//if err := sonic.ConfigStd.UnmarshalFromString(jsonStr, &m1); err != nil {
	//	logrus.Fatal(err)
	//}
	//logrus.Info("end")
}
