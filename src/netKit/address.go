// Package netKit
/*
PS: Golang对于host的定义 不同于 js.

e.g. "127.0.0.1:8888"
host in Golang:	"127.0.0.1"
host in js: 	"127.0.0.1:8888"
*/
package netKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"net"
	"strconv"
)

type (
	Address struct {
		Host string `json:"hostname,omitempty"`
		Port int    `json:"port,omitempty"`
	}
)

func (addr Address) String() string {
	return JoinHostPort(addr.Host, addr.Port)
}

// ParseStringToAddress
/*
e.g.
("https://127.0.0.1") 		=> "127.0.0.1:443", nil
("http://127.0.0.1:8888") 	=> "127.0.0.1:8888", nil
("https://blog.csdn.net/weixin_52428496/article/details/110159938") => "blog.csdn.net:443", nil
*/
func ParseStringToAddress(str string) (*Address, error) {
	str = strKit.RemoveSpace(str)
	str = strKit.ReplaceAll(str, "\\", "/")

	tmp := str
	if strKit.IsEmpty(tmp) {
		return nil, errorKit.Simple("str(\"%s\") is invalid", str)
	}
	index := strKit.Index(tmp, "://")
	if index != -1 {
		tmp = strKit.SubAfter(tmp, index+3)
	}
	index = strKit.Index(tmp, "/")
	if index != -1 {
		tmp = strKit.Substring(tmp, 0, index)
	}
	tmp = strKit.RemoveSuffixIfExist(tmp, ":")
	if strKit.IsEmpty(tmp) {
		return nil, errorKit.Simple("str(\"%s\") is invalid", str)
	}
	arr := strKit.Split(tmp, ":")
	switch len(arr) {
	case 1:
		if strKit.StartWith(str, "http://") {
			return &Address{
				Host: arr[0],
				Port: 80,
			}, nil
		} else if strKit.StartWith(str, "https://") {
			return &Address{
				Host: arr[0],
				Port: 443,
			}, nil
		}
		return nil, errorKit.Simple("str(\"%s\") is invalid", str)
	case 2:
		if port, err := intKit.ParseStringToInt(arr[1]); err != nil {
			return nil, errorKit.Simple("port string(\"%s\") is invalid", arr[1])
		} else {
			return &Address{
				Host: arr[0],
				Port: port,
			}, nil
		}
	default:
		return nil, errorKit.Simple("str(\"%s\") is invalid", str)
	}
}

// SplitHostPort 将字符串分割成 host 和 port
/*
@return 分别为: host、port、err

e.g.
("127.0.0.1:80") 	=> ("127.0.0.1", "80", nil)
("localhost:8888") 	=> ("localhost", "8888", nil)
(":80")				=> ("", 80, nil)
*/
func SplitHostPort(str string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(str)
	if err != nil {
		return "", 0, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}
	return host, port, nil
}

// JoinHostPort
/*
@param host 可以为""

e.g.
("127.0.0.1", 80) 	=> "127.0.0.1:80"
("", 80) 			=> ":80"
*/
func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}
