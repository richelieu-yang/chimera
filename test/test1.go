package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/goZeroKit"
)

func main() {
	wc, err := goZeroKit.NewSizeLimitRotateRuleWriteCloser("test.log", "-", 1, 10, 3, true)
	if err != nil {
		panic(err)
	}

	for {
		_, err := wc.Write([]byte("qwdqwdq1wdqwdqwd4684986489649864948qwdqwd强无敌群无多群无多\n"))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
