package logrusKit

import (
	"github.com/richelieu42/go-scales/src/core/file/rotateFileKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

// NewConsoleLogger 仅输出到控制台
func NewConsoleLogger(formatter logrus.Formatter, level logrus.Level) *logrus.Logger {
	if formatter == nil {
		formatter = DefaultTextFormatter
	}

	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetLevel(level)
	return logger
}

// NewFileLogger
/*
参考：https://blog.csdn.net/weixin_42681866/article/details/120876946

PS:
(1) 文件已经存在的话，会 append；
(2) 返回的 *logrus.Logger 实例，可以通过 SetLevel() 设置日志级别（即时生效；与参数 additivity 无关）.

@param logPath		会自动创建父级目录；e.g. "d:/test/test.log"
@param formatter 	可以为nil，将采用默认值
@param additivity   true: 将日志内容输出到日志文件的同时，也输出到控制台
*/
func NewFileLogger(logPath string, formatter logrus.Formatter, level logrus.Level, toConsole bool) (*logrus.Logger, error) {
	// 每多少时间生成一个新日志文件？（以防一个日志文件太大）
	rotationTime := time.Hour * 12
	// 日志文件的有效期，超时将被删除
	maxAge := timeKit.Week

	writer, err := rotateFileKit.NewRotateWriter(logPath, rotationTime, maxAge)
	if err != nil {
		return nil, err
	}
	logger := NewConsoleLogger(formatter, level)
	if toConsole {
		lfsHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.TraceLevel: writer,
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, formatter)
		logger.AddHook(lfsHook)
	} else {
		logger.Out = writer
	}
	return logger, nil
}
