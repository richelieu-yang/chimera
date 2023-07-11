package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	t1 := t.Add(-1 * timeKit.Day)
	fmt.Println(t)
	fmt.Println(t1)

	fmt.Println(timeKit.FormatTimeToString(t, "2006-01-02"))
	fmt.Println(timeKit.FormatTimeToString(t1, "2006-01-02"))

	//fmt.Println(fileKit.CutAndPaste("nohup111.out", "cyy.log"))
}
