package main

import (
	"fmt"
	"net/url"
)

func main() {
	addr := "http://localhost?wifi=true&carrier=#Staysafe AIS&os=android"

	u0, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	fmt.Println(u0)

	u1, err := url.ParseRequestURI(addr)
	if err != nil {
		panic(err)
	}
	fmt.Println(u1)
}
