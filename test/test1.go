package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
)

func main() {
	//os.File{}
	//bytes.Buffer{}
	//bytes.Reader{}
	//reader := strings.Reader
	//reader.Reset()

	//http.Request{}
	//
	//io.ReaderAt()
	//
	//bufio.Reader{}
	//
	//bufio.Writer{}
	//
	//bufio.ReadWriter{}

	//bufio.NewScanner()
	//
	//http.MaxBytesReader()

	//bufio.NewReader()
	//bufio.NewReaderSize()

	r := bytes.NewReader([]byte("0123456789"))
	b := make([]byte, 1)
	_, err := r.Read(b)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(b)) // time="2023-04-13T13:32:38+08:00" level=info msg=0

	b, err = io.ReadAll(r)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(b)) // time="2023-04-13T13:32:38+08:00" level=info msg=123456789

	// 指向最前端
	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		logrus.Fatal(err)
	}

	b, err = io.ReadAll(r)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(string(b)) // time="2023-04-13T13:32:38+08:00" level=info msg=0123456789
}
