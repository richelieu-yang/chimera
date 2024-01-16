package signalKit

import (
	"os"
	"os/signal"
)

// MonitorExitSignal 监听退出信号（拦截关闭信号）.
/*
可以参考 go-zero 中的 "proc/signals.go".

PS:
(1) 无法拦截部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）;
(2) 可以通过 logrus.RegisterExitHandler() 在程序退出前"毁尸灭迹"（在里面你甚至可以 time.Sleep）;
(3) 此函数对 主动调用os.Exit() 无效.
*/
func MonitorExitSignal(callback func(sig os.Signal)) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, ExitSignals...)

	go func() {
		sig := <-ch

		callback(sig)
		//logrus.WithField("signal", sig.String()).Fatal("Service receives an exit signal.")
	}()
}
