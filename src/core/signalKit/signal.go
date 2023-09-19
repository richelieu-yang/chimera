package signalKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var monitorOnce sync.Once

// MonitorExitSignal 监听退出信号（拦截关闭信号）.
/*
PS:
(1) 无法拦截部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）;
(2) 可以通过 logrus.RegisterExitHandler() 在程序退出前"毁尸灭迹"（在里面你甚至可以 time.Sleep）;
(3) 此函数对 主动调用os.Exit() 无效.
*/
func MonitorExitSignal() {
	monitorOnce.Do(func() {
		ch := make(chan os.Signal, 1)
		//signal.Notify(ch, exitSignals...)
		signal.Notify(ch, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)

		go func() {
			sig := <-ch
			logrus.WithField("signal", sig.String()).Fatal("Service receives an exit signal.")
		}()
	})
}
