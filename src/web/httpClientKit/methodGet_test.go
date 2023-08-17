package httpClientKit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://127.0.0.1/ping"

	logrus.Info(0)
	code, data, err := Get(url /*, WithTimeout(time.Second*3)*/)
	logrus.Info(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
	fmt.Println(data)
}
