package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t1 := t.Add(time.Second)

	fmt.Println(t.Before(t1)) // true
	fmt.Println(t.After(t1))  // false
}
