package main

import (
	"github.com/richelieu42/go-scales/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	wc, err := ioKit.NewRotateFileWriteCloser("aaa.log", time.Second*3, time.Second*30, true)
	if err != nil {
		panic(err)
	}

	wc1, err := ioKit.WrapToWriteCloser(os.Stdout)
	if err != nil {
		panic(err)
	}

	mwc, err := ioKit.MultiWriteCloser(wc, wc1)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.Out = mwc

	logger.Info("000")
	if err := mwc.Close(); err != nil {
		panic(err)
	}
	logger.Info("111")

	//if _, err := mwc.Write([]byte("0")); err != nil {
	//	panic(err)
	//}
	//if err := mwc.Close(); err != nil {
	//	panic(err)
	//}
	//if _, err := mwc.Write([]byte("1")); err != nil {
	//	panic(err)
	//}

	//if _, err := wc.Write([]byte("0")); err != nil {
	//	panic(err)
	//}
	//
	//if err := wc.Close(); err != nil {
	//	panic(err)
	//}
	//
	//if _, err := wc.Write([]byte("1")); err != nil {
	//	panic(err)
	//}

	//lock := new(sync.Mutex)
	//flag := false
	//
	//go func(wc io.WriteCloser) {
	//	time.Sleep(time.Second * 6)
	//
	//	lock.Lock()
	//	defer lock.Unlock()
	//
	//	if err := wc.Close(); err != nil {
	//		panic(err)
	//	}
	//	flag = true
	//	fmt.Println("closed")
	//}(wc)
	//
	//for {
	//	func() {
	//		lock.Lock()
	//		defer lock.Unlock()
	//
	//		if flag {
	//			fmt.Println("!!!!")
	//		}
	//		_, err := io.WriteString(wc, time.Now().String()+"\n")
	//		if err != nil {
	//			panic(err)
	//		}
	//		time.Sleep(time.Millisecond * 100)
	//	}()
	//}
}
