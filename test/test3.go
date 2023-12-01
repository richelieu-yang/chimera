package main

import (
	"fmt"
	"github.com/apache/rocketmq-clients/golang/v5/pkg/utils"
	"runtime"
	"runtime/debug"
)

func main() {
	fmt.Println(string(utils.GetMacAddress()))

	debug.PrintStack()

}

func DumpStacks() string {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	return string(buf)
}
