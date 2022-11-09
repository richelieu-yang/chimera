package scalesLogKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

type (
	ILogger interface {
		Debug(args ...interface{})

		Debugf(format string, args ...interface{})

		Info(args ...interface{})

		Infof(format string, args ...interface{})

		Warn(args ...interface{})

		Warnf(format string, args ...interface{})

		Error(args ...interface{})

		Errorf(format string, args ...interface{})

		Dispose() error
	}

	// Logger 对 logrus.Logger 进行了封装，以便于释放资源.
	Logger struct {
		// disposed 相关资源是否已经被释放？
		disposed bool
		// prefix 每条输出的前缀，可以为""
		prefix string

		rwLock       *sync.RWMutex
		logrusLogger *logrus.Logger

		// output 可以为nil
		output interface{}
		// filePath 可以为""
		filePath string
	}
)

func (l *Logger) Debug(args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	args = attachPrefixToArgs(l.prefix, args...)
	l.logrusLogger.Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	format = attachPrefixToFormat(l.prefix, format)
	l.logrusLogger.Debugf(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	args = attachPrefixToArgs(l.prefix, args...)
	l.logrusLogger.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	format = attachPrefixToFormat(l.prefix, format)
	l.logrusLogger.Infof(format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	args = attachPrefixToArgs(l.prefix, args...)
	l.logrusLogger.Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	format = attachPrefixToFormat(l.prefix, format)
	l.logrusLogger.Warnf(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	args = attachPrefixToArgs(l.prefix, args...)
	l.logrusLogger.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()

	if l.checkDisposed() {
		return
	}

	format = attachPrefixToFormat(l.prefix, format)
	l.logrusLogger.Errorf(format, args...)
}

func (l *Logger) Dispose() error {
	if l.disposed {
		// 资源已经被释放的情况: do nothing
		return nil
	}

	l.rwLock.Lock()
	defer l.rwLock.Unlock()

	if l.disposed {
		// 资源已经被释放的情况: do nothing
		return nil
	}

	defer func() {
		l.output = nil
		l.disposed = true
	}()

	switch l.output {
	case os.Stdout:
		fallthrough
	case os.Stderr:
		// Richelieu: 这两种情况不要关！
		return nil
	default:
		if c, ok := l.output.(io.Closer); ok {
			return c.Close()
		}
		return nil
	}
}

// checkDisposed
/*
@return true: 资源已经被释放（此时应当中断流程）
*/
func (l *Logger) checkDisposed() bool {
	if l.disposed {
		// 输出到控制台
		logrus.Errorf("scalesLogKit.Logger has already been disposed!")
	}
	return l.disposed
}

func attachPrefixToArgs(prefix string, args ...interface{}) []interface{} {
	if args == nil {
		return nil
	}
	arg := args[0]
	if str, ok := arg.(string); ok {
		args[0] = attachPrefixToFormat(prefix, str)
	}
	return args
}

func attachPrefixToFormat(prefix, format string) string {
	if strKit.IsEmpty(prefix) {
		return format
	}
	return prefix + format
}
