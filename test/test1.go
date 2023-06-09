package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func main() {
	fmt.Printf("%+v\n", a())

	fmt.Println("=========")

	fmt.Printf("%+v\n", b())
}

func a() error {
	return errorKit.New("123")
}

func b() error {
	str := "   "
	if err := strKit.AssertNotBlank(str, "str"); err != nil {
		return err
	}
	return nil
}
