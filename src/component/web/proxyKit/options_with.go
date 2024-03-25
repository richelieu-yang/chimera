package proxyKit

import "log"

func WithScheme(scheme string) ProxyOption {
	return func(opts *proxyOptions) {
		opts.scheme = scheme
	}
}

func WithErrorLogger(errorLogger *log.Logger) ProxyOption {
	return func(opts *proxyOptions) {
		opts.errorLogger = errorLogger
	}
}

// WithReqUrlPath
/*
PS: 如果 当前路由 和 目标路由 一致，可以不配置此项；否则必须配置.

@param reqUrlPath 	(1) 不带query;
					(2) 可以使用 ptrKit.Of() 生成 *string 实例.
*/
func WithReqUrlPath(reqUrlPath *string) ProxyOption {
	return func(opts *proxyOptions) {
		opts.reqUrlPath = reqUrlPath
	}
}

// WithOverrideQueryParams
/*
PS: WithOverrideQueryParams 和 WithExtraQueryParams 只能二选一.
*/
func WithOverrideQueryParams(overrideQueryParams map[string][]string) ProxyOption {
	return func(opts *proxyOptions) {
		opts.overrideQueryParams = overrideQueryParams
	}
}

// WithExtraQueryParams
/*
PS: WithOverrideQueryParams 和 WithExtraQueryParams 只能二选一.
*/
func WithExtraQueryParams(extraQueryParams map[string][]string) ProxyOption {
	return func(opts *proxyOptions) {
		opts.extraQueryParams = extraQueryParams
	}
}

func WithPolyfillHeaders(polyfillHeaders bool) ProxyOption {
	return func(opts *proxyOptions) {
		opts.polyfillHeaders = polyfillHeaders
	}
}
