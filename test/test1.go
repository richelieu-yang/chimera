package main

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
	}()
	wg.Wait()
	logrus.Info("------")

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
	}()
	wg.Wait()
	logrus.Info("======")
}
