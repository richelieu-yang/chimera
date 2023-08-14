package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx0, _ := context.WithTimeoutCause(context.TODO(), time.Second, errors.New("timeout"))
	ctx1, cancel := context.WithCancelCause(context.TODO())
	cancel = cancel

	//cancel(nil)
	time.Sleep(time.Second * 2)

	fmt.Println(ctx0.Err())
	fmt.Println(ctx1.Err())

	go func() {
		select {
		case <-ctx0.Done():
			fmt.Println("ctx0.Done()")
		}
	}()

	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1.Done()")
		}
	}()

	for {

	}
}
