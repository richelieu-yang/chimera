package main

import (
	"fmt"
	"os"
)

func main() {
	print(nil)
	print(os.Stderr)

}

func print(obj interface{}) {
	switch obj {
	case os.Stdout:
		fmt.Println("os.Stdout")
	case os.Stderr:
		fmt.Println("os.Stderr")
	default:
		fmt.Println("default")
	}
}
