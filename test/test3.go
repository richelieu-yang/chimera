package main

import (
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

func main() {
	api := sonic.ConfigStd
	m := map[string]interface{}{
		"0": 3.1415926,
		"1": 1,
	}
	jsonStr, err := api.MarshalToString(m)
	if err != nil {
		logrus.Fatal(err)
	}
	var m1 map[string]interface{}
	if err := api.UnmarshalFromString(jsonStr, &m1); err != nil {
		logrus.Fatal(err)
	}
}
