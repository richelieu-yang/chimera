package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	s, err := process.Pids()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(s))
}
