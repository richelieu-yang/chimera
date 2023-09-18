package gozeroKit

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/logrusx"
)

// SetLogrusWriter go-zero与第三方日志库logrus集成
/*
!!!: 需要在 rest.MustNewServer() 后调用此方法，否则会被覆盖.

@param options 可以修改 Out、Formatter...
*/
func SetLogrusWriter(options ...func(logger *logrus.Logger)) {
	writer := logrusx.NewLogrusWriter(options...)
	logx.SetWriter(writer)
}
