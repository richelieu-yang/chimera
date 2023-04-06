package logrusKit

import (
	"github.com/richelieu42/chimera/v2/core/ioKit"
	"github.com/sirupsen/logrus"
	"time"
)

// NewRotateFileLogger
/*
Deprecated: 不推荐使用.

PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, invalid argument）.

@param filePath	内部会做处理:
				(1) 会尝试创建父级目录
				(2) 文件不存在，会自动创建
				(3) 文件存在：是个文件，追加在最后；是个目录，返回error
*/
func NewRotateFileLogger(filePath string, rotationTime, maxAge time.Duration, softLinkFlag bool, formatter logrus.Formatter, level logrus.Level, toConsoleFlag bool) (*logrus.Logger, error) {
	wc, err := ioKit.NewRotateFileWriteCloser(filePath, rotationTime, maxAge, softLinkFlag)
	if err != nil {
		return nil, err
	}

	/* 此方法不方便 Close()，因为是通过Hook实现输出到控制台的同时也输出到文件日志 */
	//logger := NewLogger(formatter, level)
	//if toConsoleFlag {
	//	// (1) 输出到: 文件日志 + 控制台
	//	lfsHook := lfshook.NewHook(lfshook.WriterMap{
	//		logrus.TraceLevel: wc,
	//		logrus.DebugLevel: wc,
	//		logrus.InfoLevel:  wc,
	//		logrus.WarnLevel:  wc,
	//		logrus.ErrorLevel: wc,
	//		logrus.FatalLevel: wc,
	//		logrus.PanicLevel: wc,
	//	}, formatter)
	//	logger.AddHook(lfsHook)
	//} else {
	//	// (2) 输出到: 文件日志
	//	logger.Out = wc
	//}
	//return logger, nil

	return newFileLogger(formatter, level, wc, toConsoleFlag), nil
}
