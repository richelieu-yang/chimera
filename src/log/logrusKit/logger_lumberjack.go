package logrusKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
)

// NewLumberjackLogger
/*
Deprecated: 建议组合使用 ioKit.NewLumberjackWriteCloser() && logrusKit.NewLogger()

@param lumberjackOptions 	不能为nil（至少要配置filePath）
@param options 				可以不传（此时为nil）

e.g.
	logger, err := logrusKit.NewLumberjackLogger([]ioKit.LumberjackOption{
		ioKit.WithFilePath("a.log"),
		ioKit.WithConsole(true),
		//ioKit.WithMaxSize(1),
		ioKit.WithCompress(true),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		logger.Info(666)
	}
*/
func NewLumberjackLogger(lumberjackOptions []ioKit.LumberjackOption, options ...LoggerOption) (*logrus.Logger, error) {
	output, err := ioKit.NewLumberjackWriteCloser(lumberjackOptions...)
	if err != nil {
		return nil, err
	}

	options = append(options, WithWriter(output))
	return NewLogger(options...), nil
}
