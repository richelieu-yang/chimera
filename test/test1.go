package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func main() {
	m := map[string]interface{}{
		"a": "0",
		"b": "1",
	}

	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalIndent(m, "", "")
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(data))

	data, err = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(m)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(data))
}
