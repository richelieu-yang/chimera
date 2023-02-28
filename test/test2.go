package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

func main() {
	s := strKit.Split("123-", "-")
	fmt.Println(s)

	//u, err := urlKit.Parse("Http://localhost:8080/go?a=123&b=456")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(u.Scheme)     // http
	//fmt.Println(u.Host)       // localhost:8080
	//fmt.Println(u.Hostname()) // localhost
	//fmt.Println(u.Port())     // 8080
	//fmt.Println(u.Path)       // /go
	//fmt.Println(u.RawQuery)   // a=123&b=456
	//fmt.Println(u.Query())    // map[a:[123] b:[456]]
}
