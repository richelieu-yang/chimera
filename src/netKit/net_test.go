package netKit

import (
	"fmt"
	"testing"
)

func TestListen(t *testing.T) {
	addrress := "127.0.0.1:12345"

	{
		l, err := Listen("tcp", addrress)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer l.Close()
	}
	{
		l, err := Listen("tcp", addrress)
		if err != nil {
			// 你试图监听的端口已经被其他进程占用
			fmt.Println(err) // listen tcp 127.0.0.1:12345: bind: address already in use
			return
		}
		defer l.Close()
	}
}
