package main

import (
	"github.com/bytedance/sonic"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	defer func() {
		if obj := recover(); obj != nil {
			logrus.Infof("recover: [%T, %v].", obj, obj)
		}
	}()

	m := map[string]interface{}{
		"0": 0,
		"1": 1,
	}
	jsonStr, err := sonic.ConfigStd.MarshalToString(m)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("jsonStr: [%s].", jsonStr)
	var m1 map[string]interface{}
	if err := sonic.ConfigStd.UnmarshalFromString(jsonStr, &m1); err != nil {
		logrus.Fatal(err)
	}
}
