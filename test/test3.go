package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
)

func main() {
	//t, err := timeKit.Parse(timeKit.FormatDate, "2022-01-01")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(t)
	//fmt.Println(t.Add(-timeKit.Day))

	//fmt.Println(timeKit.ParseDuration(""))

	//d := time.Minute*63 + time.Second*15
	//fmt.Println(timeKit.FormatDuration(d)) // 1h3m15s

	//str := timeKit.Format(time.Now(), timeKit.FormatCommon)
	//fmt.Println(str) // 2023-08-14T17:10:17.057

	fmt.Println(timeKit.FormatCurrent("") == "")
}
