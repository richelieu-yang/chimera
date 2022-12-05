package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/osKit"
)

func main() {
	str := osKit.GetEnv("JAVA_HOME")
	fmt.Println(str)
}
