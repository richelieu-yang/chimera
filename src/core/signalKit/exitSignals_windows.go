package signalKit

import (
	"os"
	"syscall"
)

var (
	exitSignals = []os.Signal{syscall.SIGTERM}

	//// exitSignals 会将程序退出的所有能捕捉的信号
	///*
	//	参考: https://mp.weixin.qq.com/s/ATiAnX0PrqkBVnSFBTiMGQ
	//*/
	//exitSignals = []os.Signal{
	//	syscall.SIGHUP,
	//	//syscall.SIGUSR1,
	//	//syscall.SIGUSR2,
	//	syscall.SIGINT,
	//	syscall.SIGTERM,
	//	//syscall.SIGTSTP,
	//	syscall.SIGQUIT,
	//	//syscall.SIGSTOP,
	//	syscall.SIGKILL,
	//}
)
