package netKit

import "net"

var (
	// SplitHostnamePort
	/*
		e.g.
			fmt.Println(SplitHostnamePort("localhost:8080")) // "localhost" "8080" <nil>
			fmt.Println(SplitHostnamePort("127.0.0.1")) // "" "" address 127.0.0.1: missing port in address
			fmt.Println(SplitHostnamePort("localhost")) // "" "" address localhost: missing port in address
	*/
	SplitHostnamePort func(host string) (hostname, port string, err error) = net.SplitHostPort
)
