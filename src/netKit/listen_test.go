package netKit

import "testing"

func TestListen(t *testing.T) {
	l, err := Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer l.Close()
}
