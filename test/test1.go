package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func main() {
	m := make(map[string]string)
	if err := jsoniter.UnmarshalFromString("        ", &m); err != nil {
		logrus.Fatal(err)
	}
}
