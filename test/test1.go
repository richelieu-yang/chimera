package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/jsonKit"
)

func main() {
	jsonKit.SetRespProcessor(func(resp *jsonKit.Response) any {
		return resp
	})
	fmt.Println(jsonKit.SealFully(nil, "c", "m", "d"))
}
