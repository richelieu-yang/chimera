package base64Kit

import "encoding/base64"

type (
	base64Options struct {
		encoding *base64.Encoding
		padding  *rune
	}

	Base64Option func(opts *base64Options)
)

// Encode
/*
参考: base64.Encoding.EncodeToString()
*/
func (opts *base64Options) Encode(src []byte) []byte {
	enc := opts.encoding

	buf := make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(buf, src)
	return buf
}

// Decode
/*
参考: base64.Encoding.DecodeString()
*/
func (opts *base64Options) Decode(src []byte) ([]byte, error) {
	enc := opts.encoding

	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}

func (opts *base64Options) EncodeToString(src []byte) string {
	return opts.encoding.EncodeToString(src)
}

func (opts *base64Options) DecodeString(s string) ([]byte, error) {
	return opts.encoding.DecodeString(s)
}

func loadOptions(options ...Base64Option) *base64Options {
	opts := &base64Options{
		encoding: nil,
		padding:  nil,
	}

	for _, option := range options {
		option(opts)
	}

	if opts.encoding == nil {
		// 默认值
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
