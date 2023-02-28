package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/netKit"
)

func main() {
	fmt.Println(netKit.ParseToAddress("https://blog.csdn.net/weixin_52428496/article/details/110159938"))

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
