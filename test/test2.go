package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func main() {
	http.MaxBytesReader()

	r := bytes.NewReader([]byte("0123456789"))
	b := make([]byte, 1)
	_, err := r.Read(b)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(b))

	b, err = io.ReadAll(r)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(b))

	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		logrus.Fatal(err)
	}

	b, err = io.ReadAll(r)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(b))
}
