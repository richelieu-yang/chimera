package main

import (
	"context"
	"fmt"
)

func main() {
	// 父Context
	ctx0, cancel0 := context.WithCancel(context.TODO())
	// 子Context
	ctx1, cancel1 := context.WithCancel(ctx0)

	cancel0 = cancel0
	cancel1 = cancel1

	cancel1()

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
