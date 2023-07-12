package base64Kit

import "encoding/base64"

type (
	options struct {
		encoding *base64.Encoding
		padding  *rune
	}

	Base64Option func(opts *options)
)

func loadOptions(base64Options ...Base64Option) *options {
	opts := &options{
		encoding: nil,
		padding:  nil,
	}

	for _, option := range base64Options {
		option(opts)
	}

	if opts.encoding == nil {
		opts.encoding = base64.StdEncoding
	}
	if opts.padding != nil {
		opts.encoding = opts.encoding.WithPadding(*opts.padding)
	}
	return opts
}

func WithEncoding(encoding *base64.Encoding) Base64Option {
	return func(opts *options) {
		opts.encoding = encoding
	}
}

func WithPadding(padding *rune) Base64Option {
	return func(opts *options) {
		opts.padding = padding
	}
}
