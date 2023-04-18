package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		logrus.Fatal(err)
	}
}
