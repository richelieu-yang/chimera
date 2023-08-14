package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.TODO())
	ctx1, _ := context.WithTimeoutCause(context.TODO(), time.Second, errors.New("timeout"))
	cancel = cancel

	//cancel(nil)
	time.Sleep(time.Second * 2)

	fmt.Println(ctx1.Err())
	fmt.Println(ctx.Err())

	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1.Done()")
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done()")
		}
	}()

	for {

	}
}
