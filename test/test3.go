package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(now)
}
