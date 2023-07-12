package base64Kit

import "encoding/base64"

type (
	base64Options struct {
		encoding *base64.Encoding
		padding  *rune
	}

	Base64Option func(opts *base64Options)
)

func loadOptions(options ...Base64Option) *base64Options {
	opts := &base64Options{
		encoding: nil,
		padding:  nil,
	}

	for _, option := range options {
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

// WithEncoding
/*
@param encoding base64.StdEncoding || base64.URLEncoding || base64.RawStdEncoding || base64.RawURLEncoding
*/
func WithEncoding(encoding *base64.Encoding) Base64Option {
	return func(opts *base64Options) {
		opts.encoding = encoding
	}
}

// WithPadding
/*
PS:
(1) base64.StdEncoding 和 base64.URLEncoding 的padding:			'='（61）
(2) base64.RawStdEncoding 和 base64.RawURLEncoding 的padding:	-1（base64.NoPadding）
*/
func WithPadding(padding *rune) Base64Option {
	return func(opts *base64Options) {
		opts.padding = padding
	}
}
