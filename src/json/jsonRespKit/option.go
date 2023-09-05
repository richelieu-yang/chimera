package jsonRespKit

import (
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

type (
	options struct {
		api jsonKit.API
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
