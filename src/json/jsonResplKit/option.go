package jsonResplKit

import "github.com/bytedance/sonic"

type (
	options struct {
		api API
	}

	Option func(opts *options)
)

func WithAPI(api API) Option {
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
		opts.api = sonic.ConfigDefault
	}
	return opts
}
