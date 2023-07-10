package jsonResplKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/msgKit"
)

var (
	provider     RespProvider        = nil
	api          API                 = nil
	msgProcessor func(string) string = nil
)

func SetUp(respProvider RespProvider, msgFilePaths []string, options ...Option) error {
	if err := interfaceKit.AssertNotNil(respProvider, "respProvider"); err != nil {
		return err
	}
	if err := msgKit.ReadFiles(msgFilePaths...); err != nil {
		return err
	}

	provider = respProvider
	opts := loadOptions(options...)
	api = opts.api
	msgProcessor = opts.msgProcessor
	return nil
}
