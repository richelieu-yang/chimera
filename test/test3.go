package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.New()
	// 我们通常推荐使用 Spin
	h.Spin()
}
