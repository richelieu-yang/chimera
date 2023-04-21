package httpClientKit

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestPost(t *testing.T) {
	code, data, err := Upload("http://127.0.0.1:8080/", map[string]string{
		"file": "/Users/richelieu/Documents/ino/notes/Golang/Golang.wps",
	}, WithPostParams(map[string]string{
		"k": "vvvvvvvvvvvvvvvv",
	}))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(code)
	logrus.Info(string(data))
}
