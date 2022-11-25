package logrusKit

import (
	"github.com/richelieu42/go-scales/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger != nil {
		return ioKit.CloseWriter(logger.Out)
	}
	return nil
}

// NewLogger 输出到控制台（os.Stderr）
func NewLogger() *logrus.Logger {
	return NewCustomizedLogger(nil, logrus.DebugLevel)
}

// NewCustomizedLogger 输出到控制台（os.Stderr）
/*
@param formatter 可以为nil，此时将采用默认值
*/
func NewCustomizedLogger(formatter logrus.Formatter, level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	if formatter == nil {
		formatter = DefaultTextFormatter
	}
	logger.SetFormatter(formatter)
	logger.SetLevel(level)
	return logger
}

// NewFileLogger
/*
PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, write /Users/richelieu/Downloads/a.txt: file already closed）.

@param toConsoleFlag true: 输出到日志文件的同时，也输出到控制台
*/
func NewFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, toConsoleFlag bool) (*logrus.Logger, error) {
	var out io.WriteCloser
	out, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	if toConsoleFlag {
		tmp, err := ioKit.WrapToWriteCloser(os.Stdout)
		if err != nil {
			return nil, err
		}
		mwc, err := ioKit.MultiWriteCloser(out, tmp)
		if err != nil {
			return nil, err
		}
		out = mwc
	}

	logger := NewCustomizedLogger(formatter, level)
	logger.Out = out
	return logger, nil
}

func NewRotateFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, rotationTime, maxAge time.Duration, toConsoleFlag bool) (*logrus.Logger, error) {
	var out io.WriteCloser
	out, err := ioKit.NewRotateFileWriteCloser(filePath, rotationTime, maxAge, true, toConsoleFlag)
	if err != nil {
		return nil, err
	}
	if toConsoleFlag {
		tmp, err := ioKit.WrapToWriteCloser(os.Stdout)
		if err != nil {
			return nil, err
		}
		mwc, err := ioKit.MultiWriteCloser(out, tmp)
		if err != nil {
			return nil, err
		}
		out = mwc
	}

	logger := NewCustomizedLogger(formatter, level)
	logger.Out = out
	return logger, nil
}

//func NewFileLogger(filePath string, toConsole bool, formatter logrus.Formatter, level logrus.Level) (*logrus.Logger, error) {
//	wc, err := ioKit.NewFileWriterCloser(filePath, toConsole)
//	if err != nil {
//		return nil, err
//	}
//
//	logger := NewCustomizedLogger(formatter, level)
//	logger.Out = wc
//	return logger, nil
//}
//
//func NewRotateFileLogger(filePath string, toConsole bool, formatter logrus.Formatter, level logrus.Level) (*logrus.Logger, error) {
//	wc, err := ioKit.NewRotateFileWriteCloser(filePath, time.Hour*12, timeKit.Week, toConsole, true)
//	if err != nil {
//		return nil, err
//	}
//
//	logger := NewCustomizedLogger(formatter, level)
//	if toConsole {
//		/* (1) 输出到：文件、控制台 */
//		lfsHook := lfshook.NewHook(lfshook.WriterMap{
//			logrus.TraceLevel: wc,
//			logrus.DebugLevel: wc,
//			logrus.InfoLevel:  wc,
//			logrus.WarnLevel:  wc,
//			logrus.ErrorLevel: wc,
//			logrus.FatalLevel: wc,
//			logrus.PanicLevel: wc,
//		}, formatter)
//		logger.AddHook(lfsHook)
//	} else {
//		/* (2) 输出到：文件 */
//		logger.Out = wc
//	}
//	return logger, nil
//}

//func NewFileLogger(logPath string, formatter logrus.Formatter, level logrus.Level, toConsole bool) (*logrus.Logger, error) {
//	return NewCustomizedFileLogger(logPath, formatter, level, toConsole, -1, -1)
//}

//// NewCustomizedFileLogger
///*
//参考：
//golang日志框架--logrus+lfshook+file-rotatelogs https://blog.csdn.net/weixin_42681866/article/details/120876946
//
//PS:
//(1) 文件已经存在的话，会 append；
//(2) 返回的 *logrus.Logger 实例，可以通过 SetLevel() 设置日志级别（即时生效；与参数 additivity 无关）.
//
//@param logPath		会自动创建父级目录；e.g. "d:/test/test.log"
//@param formatter 	可以为nil，将采用默认值
//@param level 		日志级别
//@param toConsole   	将日志内容输出到日志文件的同时，是否也输出到控制台？
//@param rotationTime 每多少时间生成一个新日志文件？（以防一个日志文件太大）
//@param maxAge 		日志文件的有效期（超时将被删除）
//*/
//func NewCustomizedFileLogger(logPath string, formatter logrus.Formatter, level logrus.Level, toConsole bool, rotationTime, maxAge time.Duration) (*logrus.Logger, error) {
//	/* 处理 logPath */
//	if err := strKit.AssertNotEmpty(logPath, "logPath"); err != nil {
//		return nil, err
//	}
//	if err := fileKit.MkParentDirs(logPath); err != nil {
//		return nil, err
//	}
//
//	writer, err := rotateFileKit.NewRotateWriter(logPath, rotationTime, maxAge)
//	if err != nil {
//		return nil, err
//	}
//	logger := NewCustomizedLogger(formatter, level)
//	if toConsole {
//		lfsHook := lfshook.NewHook(lfshook.WriterMap{
//			logrus.TraceLevel: writer,
//			logrus.DebugLevel: writer,
//			logrus.InfoLevel:  writer,
//			logrus.WarnLevel:  writer,
//			logrus.ErrorLevel: writer,
//			logrus.FatalLevel: writer,
//			logrus.PanicLevel: writer,
//		}, formatter)
//		logger.AddHook(lfsHook)
//	} else {
//		logger.Out = writer
//	}
//	return logger, nil
//}
