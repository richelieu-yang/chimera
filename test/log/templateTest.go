package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"github.com/sirupsen/logrus"
)

func main() {
	m := make(map[string]interface{})
	m["test"] = 666
	jsonStr, err := jsonKit.MarshalToStringWithIndent(m)
	if err != nil {
		panic(err.Error())
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote: true,
	})
	logrus.Info(jsonStr)

	fmt.Println(logrus.TextFormatter{
		DisableQuote: false,
	}.DisableQuote)

	logrus.SetFormatter(&logrus.TextFormatter{
		//DisableQuote: false,
	})
	logrus.Info(jsonStr)
}
