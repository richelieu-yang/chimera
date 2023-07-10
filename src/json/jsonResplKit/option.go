package jsonResplKit

import "github.com/bytedance/sonic"

type (
	options struct {
		api API

		msgProcessor func(string) string
	}

	Option func(opts *options)
)

func WithAPI(api API) Option {
	return func(opts *options) {
		opts.api = api
	}
}

func WithMsgProcessor(msgProcessor func(string) string) Option {
	return func(opts *options) {
		opts.msgProcessor = msgProcessor
	}
}

func loadOptions(optionSlice ...Option) *options {
	opts := &options{
		api:          nil,
		msgProcessor: nil,
	}
	for _, option := range optionSlice {
		option(opts)
	}
	if opts.api == nil {
		opts.api = sonic.ConfigDefault
	}
	return opts
}
