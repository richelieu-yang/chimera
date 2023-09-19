package signalKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

// MonitorExitSignal 监听退出信号（拦截关闭信号）.
/*
PS:
(1) 无法拦截部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）；
(2) 可以通过 logrus.RegisterExitHandler() 在程序退出前"毁尸灭迹".
*/
func MonitorExitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, exitSignals...)

	go func() {
		sig := <-ch
		logrus.WithField("signal", sig.String()).Fatal("Service receives an exit signal.")

		//logrus.WithField("signal", sig.String()).Warn("receive a signal and this process will exit")
		//if exitFunc != nil {
		//	logrus.Debug("exitFunc starts")
		//	exitFunc(sig)
		//	logrus.Debug("exitFunc ends")
		//}
		//os.Exit(0)
	}()
}
