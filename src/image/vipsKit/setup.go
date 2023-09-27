package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/sirupsen/logrus"
)

// SetUp
/*
@param config 可以为nil（使用默认配置: concurrency=1 cache_max_files=0 cache_max_mem=52428800 cache_max=100）
*/
func SetUp(config *vips.Config) {
	vips.Startup(config)

	logrus.RegisterExitHandler(func() {
		vips.Shutdown()
	})
}
