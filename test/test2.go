package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/urlKit"
)

func main() {
	u, err := urlKit.Parse("http://localhost:8080/go?a=123&b=456")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.RawQuery) // a=123&b=456

	m, err := urlKit.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(m) // map[a:[123] b:[456]]
}
