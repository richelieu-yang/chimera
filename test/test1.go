package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan error, 1)

	go func() {
		time.Sleep(time.Second)
		ch <- nil
	}()

	select {
	case err := <-ch:
		fmt.Println(err)
	}
}
