// Package netKit
/*
PS: Golang对于host的定义 不同于 js.

e.g. "127.0.0.1:8888"
host in Golang:	"127.0.0.1"
host in js: 	"127.0.0.1:8888"
*/
package netKit

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"net"
	"strconv"
)

type (
	Address struct {
		// Hostname 也可以是ip
		Hostname string
		Port     int
	}
)

func (addr *Address) String() string {
	return fmt.Sprintf("%s:%d", addr.Hostname, addr.Port)
}

func JoinHostnameAndPort(hostname string, port int) string {
	//addr := &Address{
	//	Hostname: hostname,
	//	Port:     port,
	//}
	//return addr.String()

	return net.JoinHostPort(hostname, strconv.Itoa(port))
}

// ParseToAddress
/*
e.g.
("https://127.0.0.1") 		=> "127.0.0.1:443", nil
("http://127.0.0.1:8888") 	=> "127.0.0.1:8888", nil
("https://blog.csdn.net/weixin_52428496/article/details/110159938") => "blog.csdn.net:443", nil
*/
func ParseToAddress(str string) (*Address, error) {
	tmp := strKit.Trim(str)
	tmp = strKit.ReplaceAll(tmp, "\\", "/")
	if strKit.IsEmpty(tmp) {
		return nil, errorKit.Simple("str(\"%s\") is invalid", str)
	}

	var scheme string
	index := strKit.Index(tmp, "://")
	if index != -1 {
		scheme = strKit.SubBefore(tmp, index)
		scheme = strKit.ToLower(scheme)
		tmp = strKit.SubAfter(tmp, index+3)
	}
	index = strKit.Index(tmp, "/")
	if index != -1 {
		tmp = strKit.Substring(tmp, 0, index)
	}
	if strKit.IsEmpty(tmp) {
		return nil, errorKit.Simple("invalid str(%s)", str)
	}

	arr := strKit.Split(tmp, ":")
	switch len(arr) {
	case 1:
		hostname := arr[0]
		port, err := getDefaultPortByScheme(scheme)
		if err != nil {
			return nil, err
		}
		return &Address{
			Hostname: hostname,
			Port:     port,
		}, nil
	case 2:
		hostname := arr[0]
		portStr := arr[1]
		if strKit.IsAllEmpty(hostname, portStr) {
			return nil, errorKit.Simple("invalid str(%s)", str)
		}

		var port int
		var err error
		if strKit.IsEmpty(portStr) {
			port, err = getDefaultPortByScheme(scheme)
		} else {
			port, err = strconv.Atoi(portStr)
		}
		if err != nil {
			return nil, err
		}

		return &Address{
			Hostname: hostname,
			Port:     port,
		}, nil
	default:
		return nil, errorKit.Simple("invalid str(%s)", str)
	}
}

func getDefaultPortByScheme(scheme string) (int, error) {
	switch scheme {
	case "http":
		return 80, nil
	case "https":
		return 443, nil
	default:
		return -1, errorKit.Simple("invalid scheme(%s)", scheme)
	}
}
