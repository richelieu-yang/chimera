// Package netKit
/*
PS: Golang对于host的定义 不同于 js.

e.g. "127.0.0.1:8888"
host in Golang:	"127.0.0.1"
host in js: 	"127.0.0.1:8888"
*/
package netKit

type (
	Address struct {
		Host string
		Port int
	}
)

func (addr Address) String() string {
	return JoinHostPort(addr.Host, addr.Port)
}
