package jsonRespKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
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

	for _, path := range opts.filePaths {
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
