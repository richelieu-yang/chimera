package main

import (
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	pool, err := ants.NewPool(100)
	if err != nil {
		panic(err)
	}

	err = pool.Submit(func() {
		panic(666)
	})
	if err != nil {
		panic(err)
	}

	logrus.Info("---")
	select {}
}
