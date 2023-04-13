package main

import (
	"bytes"
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func main() {
	bytes.NewReader()

	reader := strings.NewReader("test")
	readCloser := ioKit.NopCloserToReader(reader)
	readCloser = http.MaxBytesReader(nil, readCloser, 2)

	data := make([]byte, 3)
	_, err := readCloser.Read(data)
	if err != nil {
		if mbe, ok := err.(*http.MaxBytesError); ok {
			logrus.Infof("*http.MaxBytesError: %s", mbe.Error())
		} else {
			logrus.Fatal(err)
		}
	}
}
