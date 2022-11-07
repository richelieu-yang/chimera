package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/http/refererKit"
)

func main() {
	v, err := refererKit.NewRefererVerifier(true, false, "", "*.yozo.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(v.Verify("", ""))
	fmt.Println(v.Verify("", "http://www.yozo.com"))
	fmt.Println(v.Verify("", "http://mail.yozo.com"))
	fmt.Println(v.Verify("", "http://1.yozo.com1"))

	//re, err := regexpKit.StringToRegexp("**.yozo.com")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(re)
	//fmt.Println(re.MatchString("11yozo2com"))
	//fmt.Println(re.MatchString(".yozo.com"))
	//fmt.Println(re.MatchString("1.yozo.com"))
	//fmt.Println(re.MatchString("www.yozo.com"))
}
