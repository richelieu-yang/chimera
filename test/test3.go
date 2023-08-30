package main

import (
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	pid := 100
	p, err := process.NewProcess(pid)
	if err != nil {
		panic(err)
	}
	p.CPUPercent()
}
