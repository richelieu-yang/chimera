package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	perm := os.FileMode(0644)
	if perm&os.ModePerm == os.ModePerm {
		// 文件有写权限
	}

	os.IsSupport256Color()

	sig := syscall.SIGTERM
	fmt.Printf("%v\n%s\n%s\n", sig, sig, sig.String())
}
