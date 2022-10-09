package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/jsonKit"
)

type (
	CurJsonResponse struct {
		Code    string      `json:"errorCode" example:"0"`
		Message string      `json:"errorMessage" example:"no error"`
		Data    interface{} `json:"result,omitempty"`
	}
)

func main() {
	jsonKit.SetJsonResponseProcessor(func(resp *jsonKit.JsonResponse) any {
		return &CurJsonResponse{
			Code:    resp.Code,
			Message: resp.Message,
			Data:    resp.Data,
		}
	})

	fmt.Println(jsonKit.SealFully(nil, "c", "m", "d"))
}
