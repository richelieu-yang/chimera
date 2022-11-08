package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/regexpKit"
)

func main() {
	re, err := regexpKit.StringToRegexp("")
	if err != nil {
		panic(err)
	}

	fmt.Println(re.MatchString(""))
	fmt.Println(re.MatchString("awdqw"))
	fmt.Println(re.MatchString("强无敌群无多"))

}
