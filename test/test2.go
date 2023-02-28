package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/urlKit"
)

func main() {
	u, err := urlKit.Parse("localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)     // http
	fmt.Println(u.Host)       // localhost:8080
	fmt.Println(u.Hostname()) // localhost
	fmt.Println(u.Port())     // 8080
	fmt.Println(u.Path)       // /go
	fmt.Println(u.RawQuery)   // a=123&b=456
	fmt.Println(u.Query())    // map[a:[123] b:[456]]
}
