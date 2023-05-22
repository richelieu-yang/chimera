package jsonKit

import jsoniter "github.com/json-iterator/go"

type (
	jsonOptions struct {
		api jsoniter.API
		// 目前只能是 ""
		prefix string
		/*
			(1) 目前只能是 "" 或 多个空格（不能有其他字符）
			(2) encoding/json标准库 可以用"\t"
		*/
		indent string
	}

	JsonOption func(opts *jsonOptions)
)

func loadOptions(options ...JsonOption) *jsonOptions {
	opts := &jsonOptions{
		api:    nil,
		prefix: "",
		indent: "",
	}

	for _, option := range options {
		option(opts)
	}

	// 容错，以防调用方瞎搞
	if opts.api == nil {
		opts.api = jsoniter.ConfigDefault
	}

	return opts
}

// WithApi
/*
@param api jsoniter.ConfigDefault || jsoniter.ConfigCompatibleWithStandardLibrary || jsoniter.ConfigFastest
*/
func WithApi(api jsoniter.API) JsonOption {
	return func(opts *jsonOptions) {
		opts.api = api
	}
}

func WithIndent(indent string) JsonOption {
	return func(opts *jsonOptions) {
		opts.indent = indent
	}
}
