package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	filename := "./test.log"
	logger, err1 := logx.NewLogger(
		filename,
		logx.DefaultRotateRule(
			filename,
			"-",
			1,
			true,
		),
		true,
	)
	if err1 != nil {
		panic(err1)
	}
	defer logger.Close()

	fmt.Println(logger.Write([]byte("cyy\n")))
}
