package httpClientKit

import "time"

const (
	// DefaultTimeout 发送请求的默认超时时间.
	/*
		PS:
		(1) 个人实测，对于 http.Client 结构体，Timeout 默认为30s.
		(2) e.g. yozo的网访问谷歌必定超时.
	*/
	DefaultTimeout = time.Second * 10
)

type (
	options struct {
		timeout time.Duration
		// safe 默认false（即使客户的url以https开头&&证书非法，请求也能成功）
		safe        bool
		queryParams map[string]string
		postParams  map[string]string
	}

	Option func(opts *options)
)

func loadOptions(s ...Option) *options {
	opts := &options{}

	for _, option := range s {
		option(opts)
	}

	if opts.timeout <= 0 {
		opts.timeout = DefaultTimeout
	}

	return opts
}

// WithTimeout
/*
适用于: GET、POST
*/
func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

// WithSafe
/*
适用于: GET、POST
*/
func WithSafe(safe bool) Option {
	return func(opts *options) {
		opts.safe = safe
	}
}

// WithQueryParams
/*
适用于: GET、POST
*/
func WithQueryParams(queryParams map[string]string) Option {
	return func(opts *options) {
		opts.queryParams = queryParams
	}
}

// WithPostParams
/*
适用于: POST
*/
func WithPostParams(postParams map[string]string) Option {
	return func(opts *options) {
		opts.postParams = postParams
	}
}
