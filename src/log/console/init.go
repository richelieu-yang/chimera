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

	// core 4个logger共享的唯一core，同时也是 Sync() 的唯一落点.
	/*
		默认的输出目标是 os.Stdout（无缓冲），详见 Sync() 的说明.
	*/
	core zapcore.Core

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
	ws := zapKit.NewLockedWriteSyncer(os.Stdout)
	core = zapKit.NewCore(encoder, ws, level)

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

// Sync 刷新底层 core 的输出缓冲.
/*
!!!: 默认的输出目标(os.Stdout)是 无缓冲 的，所以本方法在默认配置下 不会改变任何结果.

"无缓冲"意味着: 每次 Write 都会立刻发起一次 write(2) 系统调用，日志在 Write 返回时
就已经进入操作系统内核，不存在"滞留在用户态缓冲区、还没写出去"的数据.
因此本方法唯一可能做的事只是把内核页缓存刷到物理磁盘(fsync)，而这与
"进程退出时会不会丢日志"无关 —— 正常退出/panic/os.Exit 都不会丢弃已进入内核的页缓存，
只有 断电 / 内核崩溃 才会.

而且对 os.Stdout 调用同步在多数目标上根本不成立（实测 macOS/Apple M1 Pro）:
	终端(TTY)     => "inappropriate ioctl for device" (ENOTTY)
	管道          => "bad file descriptor"           (EBADF)
	/dev/null     => "operation not supported by device" (ENODEV)
	重定向到文件    => nil，此时才是一次真正的 fsync

@return 底层 core 的 Sync 结果; 上述平台相关的错误属 预期行为，调用方可以直接忽略.
*/
func Sync() error {
	return core.Sync()
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
