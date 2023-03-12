package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/http/httpKit"
)

func main() {
	str := httpKit.GetContentType([]byte(nil))
	fmt.Println(str)

	//f, err := os.Create("c.log")
	//if err != nil {
	//	panic(err)
	//}
	//syscall.CloseOnExec(int(f.Fd()))
}
