package jsonRespKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

type (
	// RespProvider
	/*
		@return (1) 返回值是一个结构体实例指针
				(2) 结构体建议加上json tag
	*/
	RespProvider func(code, msg string, data interface{}) interface{}
)

var (
	provider RespProvider = nil
)

func MustSetUp(respProvider RespProvider, options ...Option) {
	if err := SetUp(respProvider, options...); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(respProvider RespProvider, options ...Option) (err error) {
	defer func() {
		if err != nil {
			provider = nil
			msgMap = make(map[string]string)
		}
	}()

	err = interfaceKit.AssertNotNil(respProvider, "respProvider")
	if err != nil {
		return err
	}
	provider = respProvider

	opts := loadOptions(options...)
	for _, path := range opts.filePathSlice {
		if err = readFile(path); err != nil {
			return err
		}
	}
	for _, fd := range opts.fileDataSlice {
		if err = readFileData(fd); err != nil {
			return err
		}
	}
	return nil
}
