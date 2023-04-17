package signalKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

// MonitorExitSignal 监听退出信号（拦截关闭信号）
/*
缺陷: 部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）无法被拦截.

@param exitFunc (1) 可以为nil
				(2) 一般用于清理痕迹（毁尸灭迹）
*/
func MonitorExitSignal(exitFunc func(os.Signal)) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, exitSignals...)

	go func() {
		sig := <-ch

		logrus.WithField("signal", sig.String()).Warn("receive a signal and this process will exit")
		if exitFunc != nil {
			logrus.Debug("exitFunc starts")
			exitFunc(sig)
			logrus.Debug("exitFunc ends")
		}
		os.Exit(0)
	}()
}
