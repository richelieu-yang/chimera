package main

import (
	"fmt"
	"time"
)

func main() {
	//t := time.Now()
	//fmt.Println(humanize.Time(t))

	// Add 时间相加
	now := time.Now()

	// 10分钟前
	m, _ := time.ParseDuration("-10m")
	m1 := now.Add(m)
	fmt.Println(m1)

}
