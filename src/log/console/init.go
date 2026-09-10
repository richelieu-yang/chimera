package console

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// level 全局日志级别: DEBUG.
	/*
		此处用 zap.AtomicLevel 而非 zapcore.Level，原因:
		(1) 修改级别时无需重建 logger，故下面4个logger指针终身稳定;
		(2) 读写级别都是原子操作，调用方缓存过 GetL()/GetSL() 的结果也不会失效;
		(3) zap.AtomicLevel 实现了 zapcore.LevelEnabler，可以直接作为 core 的级别开关.
	*/
	level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// 下面4个logger仅在 init() 中创建一次，此后不再被修改
	l       *zap.Logger
	sl      *zap.SugaredLogger
	innerL  *zap.Logger
	innerSL *zap.SugaredLogger
)

func init() {
	initialize()
}

// initialize
/*
仅由 init() 调用，故无需加锁:
(1) go 会保证: 包的初始化（变量初始化 + init()）happens-before 任何使用方的代码;
(2) 本函数执行完毕后，4个logger指针不再被写入，因此读取它们无需任何同步.
*/
func initialize() {
	encoder := zapKit.NewEncoder()
	ws := os.Stdout
	core := zapKit.NewCore(encoder, ws, level)

	l = zapKit.NewLogger(core, zapKit.WithCallerSkip(0))
	sl = l.Sugar()
	innerL = zapKit.NewLogger(core, zapKit.WithCallerSkip(1))
	innerSL = innerL.Sugar()
}

// GetL 供外部使用（skip为0）
/*
PS: 返回值终身有效，可以缓存在包级变量中.
*/
func GetL() *zap.Logger {
	return l
}

// GetSL 供外部使用（skip为0）
/*
PS: 返回值终身有效，可以缓存在包级变量中.
*/
func GetSL() *zap.SugaredLogger {
	return sl
}

// getInnerL 供内部使用（skip为1）
func getInnerL() *zap.Logger {
	return innerL
}

// getInnerSL 供内部使用（skip为1）
func getInnerSL() *zap.SugaredLogger {
	return innerSL
}

// Sync 刷新所有logger的缓冲区
func Sync() {
	_ = l.Sync()
	_ = sl.Sync()
	_ = innerL.Sync()
	_ = innerSL.Sync()
}

// SetLogLevel 修改全局日志级别
/*
PS:
(1) 修改后立即对所有logger生效;
(2) 无需重建logger，所以不影响调用方已经取出的logger指针.
*/
func SetLogLevel(lv zapcore.Level) {
	level.SetLevel(lv)
}
