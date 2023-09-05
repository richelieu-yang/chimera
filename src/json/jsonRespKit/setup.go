package jsonRespKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

type (
	// RespProvider
	/*
		@return 建议是个结构体实例 && 加上json tag
	*/
	RespProvider func(code, msg string, data interface{}) interface{}
)

var (
	provider RespProvider = nil

	// api 用于序列化json
	api jsonKit.API = nil
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
			api = nil
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
		if err := readFile(path); err != nil {
			return err
		}
	}
	for _, fd := range opts.fileDataSlice {
		if err := readFileData(fd); err != nil {
			return err
		}
	}
	api = opts.api
	return nil
}
