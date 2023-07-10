package jsonResplKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var (
	provider RespProvider = nil

	// api 用于序列化json
	api API = nil

	// msgProcessor 供外部对最终message进行二开
	msgProcessor func(string) string = nil
)

func MustSetUp(respProvider RespProvider, msgFilePaths []string, options ...Option) {
	if err := SetUp(respProvider, msgFilePaths, options...); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(respProvider RespProvider, msgFilePaths []string, options ...Option) error {
	if err := interfaceKit.AssertNotNil(respProvider, "respProvider"); err != nil {
		return err
	}
	if err := readFiles(msgFilePaths...); err != nil {
		return err
	}

	provider = respProvider
	opts := loadOptions(options...)
	api = opts.api
	return nil
}
