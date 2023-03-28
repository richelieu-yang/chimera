package main

import (
	"github.com/richelieu42/chimera/src/gozeroKit"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	logx.Info("---")
	gozeroKit.SetLogrusWriter(nil)
	logx.Info("+++")
}
