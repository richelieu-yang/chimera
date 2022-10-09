package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/jsonKit"
)

func main() {
	jsonKit.SetRespProcessor(func(resp *jsonKit.Response) any {
		return resp
		//return &CurJsonResponse{
		//	Code:    resp.Code,
		//	Message: resp.Message,
		//	Data:    resp.Data,
		//}
	})

	fmt.Println(jsonKit.SealFully(nil, "c", "m", "d"))
}
