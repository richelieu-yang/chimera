package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/signalKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
)

func init() {
	logrusKit.MustSetUp(nil)
	signalKit.MonitorExitSignal()
}

func main() {

	//logrus.RegisterExitHandler(func() {
	//	fmt.Println("0")
	//	time.Sleep(time.Second * 3)
	//	fmt.Println("1")
	//})
	//logrus.RegisterExitHandler(func() {
	//	fmt.Println("2")
	//	time.Sleep(time.Second * 10)
	//	fmt.Println("3")
	//})
	//
	//fmt.Println("PID", processKit.PID)
	//
	////time.Sleep(time.Second)
	////os.Exit(1)
	//for {
	//}
}
