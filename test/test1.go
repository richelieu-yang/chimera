package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/regexpKit"
)

func main() {
	re, err := regexpKit.StringToRegexp("**.yozo.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(re)
	fmt.Println(re.MatchString("11yozo2com"))
	fmt.Println(re.MatchString(".yozo.com"))
	fmt.Println(re.MatchString("1.yozo.com"))
	fmt.Println(re.MatchString("www.yozo.com"))
}
