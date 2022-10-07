package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/interfaceKit"
)

type Void struct {
}

func main() {
	fmt.Println(interfaceKit.IsPtr(nil))       // false
	fmt.Println(interfaceKit.IsPtr(Void{}))    // false
	fmt.Println(interfaceKit.IsPtr(&Void{}))   // true
	fmt.Println(interfaceKit.IsPtr(new(Void))) // true
}
