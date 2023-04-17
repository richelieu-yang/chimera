package osKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// MonitorExitSignal 监听退出信号（拦截关闭信号）
/*
@param exitFunc (1) 可以为nil
				(2)
*/
func MonitorExitSignal(exitFunc func(os.Signal)) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-c

		logrus.WithField("signal", sig.String()).Warn("receive a signal and this process will exit")
		if exitFunc != nil {
			exitFunc(sig)
		}
		os.Exit(0)
	}()
}
