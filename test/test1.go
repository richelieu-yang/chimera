package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"io"
)

func main() {
	err := io.EOF
	err = errorKit.WithLocationInfo(err)

	// test/test1.go:11|main: EOF
	fmt.Printf("%v\n", err)
	// EOF
	// test/test1.go:11|main
	fmt.Printf("%+v\n", err)

}
