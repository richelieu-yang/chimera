package osKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// MonitorExitSignal 监听退出信号（拦截关闭信号）
/*
PS: 收到退出信号后，程序会退出.
*/
func MonitorExitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-c
		logrus.Fatalf("signal: [%s]", sig.String())
	}()
}
