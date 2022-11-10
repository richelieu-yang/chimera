package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/ioKit"
	"io"
	"os"
)

func main() {
	writeCloser, err := ioKit.WrapToWriteCloser(os.Stderr)
	if err != nil {
		panic(err)
	}

	fmt.Println(io.WriteString(writeCloser, "0\n"))
	fmt.Println(io.WriteString(writeCloser, "1\n"))

	fmt.Println("error:", writeCloser.Close())

	fmt.Println(io.WriteString(writeCloser, "2\n"))
}

//func InitLog() {
//	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
//	log.SetFormatter(&log.TextFormatter{})
//	log.SetOutput(os.Stdout)
//	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
//	file, err := os.OpenFile("checkemstools.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	writers := []io.Writer{
//		file,
//		os.Stdout}
//	//同时写文件和屏幕
//	fileAndStdoutWriter := io.MultiWriter(writers...)
//	if err == nil {
//		log.SetOutput(fileAndStdoutWriter)
//	} else {
//		log.Info("failed to log to file.")
//	}
//	//设置最低loglevel
//	log.SetLevel(log.InfoLevel)
//}
