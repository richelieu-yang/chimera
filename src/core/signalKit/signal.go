package signalKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
)

var monitorOnce sync.Once

// MonitorExitSignal 监听退出信号（拦截关闭信号）.
/*
PS:
(1) 无法拦截部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）；
(2) 可以通过 logrus.RegisterExitHandler() 在程序退出前"毁尸灭迹"（在里面你甚至可以 time.Sleep）.
*/
func MonitorExitSignal() {
	monitorOnce.Do(func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, exitSignals...)

		go func() {
			sig := <-ch
			logrus.WithField("signal", sig.String()).Fatal("Service receives an exit signal.")
		}()
	})
}
