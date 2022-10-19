package main

import (
	"github.com/richelieu42/go-scales/src/core/file/rotateFileKit"
	"github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

func main() {
	writer, err := rotateFileKit.NewRotateWriter("/Users/richelieu/Downloads/tmp.log", time.Second*10, time.Second*10*6, true)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("java", "-jar", "/Users/richelieu/Downloads/scales.jar")
	cmd.Stdout = writer
	cmd.Stderr = writer
	// 执行命令
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	// 等待命令结束
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	logrus.Info("-----------------------")
}
