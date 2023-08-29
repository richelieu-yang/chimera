package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %T %v", r, r)
		}
	}()

	api := sonic.ConfigStd
	m := map[string]interface{}{
		"0": 3.1415926,
		"1": 1,
	}
	jsonStr, err := api.MarshalToString(m)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("jsonStr:", jsonStr)
	var m1 map[string]interface{}
	if err := api.UnmarshalFromString(jsonStr, &m1); err != nil {
		logrus.Fatal(err)
	}
}
