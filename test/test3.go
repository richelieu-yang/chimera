package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	defer func() {
		if obj := recover(); obj != nil {
			fmt.Println(obj)
			//debug.PrintStack()
			debug.Stack()
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		panic(1)
	}()

	for {
	}
}
