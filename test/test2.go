package main

import (
	"context"
	"fmt"
)

func main() {
	_, cancel := context.WithCancel(context.TODO())
	cancel()
	cancel()
	fmt.Println("======")
}
