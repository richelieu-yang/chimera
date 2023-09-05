package jsonRespKit

import (
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

type (
	options struct {
		api jsonKit.API

		filePathSlice []string
		fileDataSlice [][]byte
	}

	Option func(opts *options)
)

// WithAPI
/*
@param api sonic.ConfigDefault（默认; 推荐） || jsoniter.ConfigDefault
*/
func WithAPI(api jsonKit.API) Option {
	return func(opts *options) {
		opts.api = api
	}
}

func WithFilePathSlice(filePathSlice []string) Option {
	return func(opts *options) {
		opts.filePathSlice = filePathSlice
	}
}

func WithFileDataSlice(fileDataSlice [][]byte) Option {
	return func(opts *options) {
		opts.fileDataSlice = fileDataSlice
	}
}

func loadOptions(optionSlice ...Option) *options {
	opts := &options{
		api: nil,
	}
	for _, option := range optionSlice {
		option(opts)
	}
	if opts.api == nil {
		opts.api = jsonKit.GetDefaultApi()
	}
	return opts
}
