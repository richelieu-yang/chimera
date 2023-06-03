package main

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func main() {
	logrusKit.MustSetUp(nil)

	_, err1 := gjson.Encode(func() {})
	err2 := errorKit.Wrap(err1, `error occurred`)
	//fmt.Printf("%+v", err2)

	//logrusKit.DisableQuoteTemporarily()

	logrus.Fatalf("%+v", err2)

	//logrus.Fatal(err2)
}
