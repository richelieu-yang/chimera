package httpKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"log"
	"net/http"
	"net/http/httputil"
)

type (
	proxyOptions struct {
		scheme      string
		errorLogger *log.Logger
		// reqUrlPath 请求路由
		reqUrlPath  *string
		queryParams map[string][]string
	}

	ProxyOption func(opts *proxyOptions)
)

func loadOptions(options ...ProxyOption) *proxyOptions {
	opts := &proxyOptions{}
	for _, option := range options {
		option(opts)
	}

	opts.scheme = strKit.EmptyToDefault(opts.scheme, "http", true)

	return opts
}

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

func WithQueryParams(queryParams map[string][]string) ProxyOption {
	return func(opts *proxyOptions) {
		opts.queryParams = queryParams
	}
}

// Proxy 代理请求（反向代理，请求转发）.
/*
PS: 转发请求前如果想变更请求头(Request Header)，可以在调用此函数前设置请求头.

@param w		e.g.ctx.Writer
@param r 		e.g.ctx.Request
@param options 	可以不传
*/
func Proxy(w http.ResponseWriter, r *http.Request, targetAddr string, options ...ProxyOption) error {
	opts := loadOptions(options...)
	return opts.proxy(w, r, targetAddr)
}

// proxy
/*
@param errLogger 	可以为nil，但不建议这么干，因为错误会输出到控制台（通过 log.Printf()），不利于错误定位
@param scheme 		"http" || "https"
@param addr 		e.g."127.0.0.1:8888"
@param reqUrlPath 	(1) 可以为nil（此时不修改 req.URL.Path）
					(2) 非nil的话，个人感觉: 字符串的第一个字符应该是"/"
@param queryParams 	可以为nil
@return 可能是 context.Canceled（可以用 == 进行比较）

更多可参考:
httputil.NewSingleHostReverseProxy() https://m.bilibili.com/video/BV1H64y1u7D7?buvid=Y44D4D448DC195994A5A88CED2DA982C60DF&is_story_h5=false&mid=5%2BiuUUrTqJQOdIa1r3VR0g%3D%3D&p=1&plat_id=114&share_from=ugc&share_medium=iphone&share_plat=ios&share_session_id=8B36D2C9-4DCB-4BE5-80AD-F7D49E292B5F&share_source=WEIXIN&share_tag=s_i&timestamp=1680160438&unique_k=16bK0gz&up_id=456307879

PS:
(1) 支持代理的协议: https、http、wss、ws...
(2) 如果请求转发的目标有效，但处理此请求需要花费大量时间（比如20+min），此时如果请求的客户端终端了请求（e.g.浏览器页面被直接关闭了），将返回 context.Canceled.
(3) addr有效，reqUrlPath非nil但事实上不存在该路由的情况，返回值为nil && 原始客户端得到404（404 page not found）.
(4) 代理请求前，如果读取了Request.Body的内容但不恢复（即重置其内容），将直接返回error（e.g.net/http: HTTP/1.x transport connection broken: http: ContentLength=161 with Body length 0）.

	且目标方不会收到请求.（TODO: 很奇怪，会走两遍，第二次返回的error为: context canceled）

e.g.	将 https://127.0.0.1:8888/test 转发给 http://127.0.0.1:8889/test
传参可以是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=nil
(2) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"/test"
传参不能是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"test" （400 Bad Request）

e.g.1	将 https://127.0.0.1:8888/test 转发给 http://127.0.0.1:8889/test1
传参可以是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"/test1"
传参不能是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"test1"

e.g.2	将 https://127.0.0.1:8888/group/test 转发给 http://127.0.0.1:8889/test1
scheme="http" addr="127.0.0.1:8889" reqUrlPath=ptrKit.ToPtr("/test1")

e.g.3	将 https://127.0.0.1:8888/group/test 转发给 http://127.0.0.1:8889/group1/test1
scheme="http" addr="127.0.0.1:8889" reqUrlPath=ptrKit.ToPtr("/group1/test1")

e.g.4	将 wss://127.0.0.1:8888/test 转发给 ws://127.0.0.1:80/ws/connect
scheme="http" addr="127.0.0.1:80" reqUrlPath=ptrKit.ToPtr("/ws/connect")
*/
func (opts *proxyOptions) proxy(w http.ResponseWriter, r *http.Request, addr string) error {
	// reset Request.Body
	if err := ResetRequestBody(r); err != nil {
		return err
	}

	//// scheme
	//scheme = strKit.EmptyToDefault(scheme, "http", true)
	//switch scheme {
	//case "https":
	//case "http":
	//default:
	//	return errorKit.New("invalid scheme: %s", scheme)
	//}

	// addr
	if err := strKit.AssertNotEmpty(addr, "addr"); err != nil {
		return err
	}

	var err error
	director := func(req *http.Request) {
		req.URL.Scheme = opts.scheme
		req.URL.Host = addr
		if opts.reqUrlPath != nil {
			req.URL.Path = *opts.reqUrlPath
		}

		// 可能会修改 req.URL.RawQuery
		urlKit.AddQueryParamsToRawQuery(req.URL, opts.queryParams)
	}
	reverseProxy := &httputil.ReverseProxy{
		Director: director,
		ErrorLog: opts.errorLogger,
		ErrorHandler: func(rw http.ResponseWriter, req *http.Request, e error) {
			err = e
		},
	}
	reverseProxy.ServeHTTP(w, r)
	return err
}
