package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	secs := t.Unix()
	t1 := time.Unix(secs, 0)

	fmt.Println(t)  // 2021-07-31 13:22:25.221929 +0800 CST m=+0.004342301
	fmt.Println(t1) // 2021-07-31 13:22:25 +0800 CST
}
