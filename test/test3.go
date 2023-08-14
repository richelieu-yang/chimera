package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
)

func main() {
	t, err := timeKit.ParseTimeString(string(timeKit.FormatDate), "2022-01-01")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println(t.Add(-timeKit.Day))
}
