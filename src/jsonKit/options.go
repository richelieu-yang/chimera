package jsonKit

import jsoniter "github.com/json-iterator/go"

type (
	jsonOptions struct {
		api jsoniter.API
	}

	JsonOption func(opts *jsonOptions)
)

func loadOptions(options ...JsonOption) *jsonOptions {
	opts := &jsonOptions{
		api: nil,
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
