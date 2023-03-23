package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/src/core/mapKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/http/httpKit"
	"github.com/richelieu42/chimera/src/netKit"
	"github.com/richelieu42/chimera/src/urlKit"
	"log"
	"net/http"
	"net/http/httputil"
)

type (
	// ForwardParams 请求转发所需的参数
	ForwardParams struct {
		netKit.Address

		// 一般为："http"（包括ws协议）、"https"
		Scheme string
		// 可以为 nil，也可以通过 GetReqUrlPathForForward() 获取
		ReqUrlPath *string
		// 拼成字符串，放在url中
		ExtraQuery map[string]string
	}
)

var baseExtraQuery map[string]string

func SetBaseExtraQuery(m map[string]string) {
	baseExtraQuery = m
}

// ForwardRequest 请求转发（http请求、websocket请求...）
/*
！！！：如果两个地址，一个有contextPath(""和"/"等价)一个没有，需要注意参数path；其他情况参数path直接传nil即可.

@param errLogger 可以为nil，但不建议这么干，因为错误会输出到控制台（通过 log.Printf()），不利于错误定位
@return 可能是 context.Canceled（可以用==进行比较）

e.g.
如果请求转发的目标有效，但处理此请求需要花费大量时间（比如20+min），此时如果请求的客户端终端了请求（e.g.浏览器页面被直接关闭了），将返回 context.Canceled.
*/
func ForwardRequest(ctx *gin.Context, params ForwardParams, errorLogger *log.Logger) error {
	var err error

	// scheme默认"http"
	scheme := strKit.EmptyToDefault(params.Scheme, "http", true)
	addr := params.Address.String()
	reqUrlPath := params.ReqUrlPath
	extraQuery := mapKit.Merge(baseExtraQuery, params.ExtraQuery)

	director := func(req *http.Request) {
		req.URL.Scheme = scheme
		req.URL.Host = addr

		// req.URL.Param1
		if reqUrlPath != nil {
			req.URL.Path = *reqUrlPath
		}

		// req.URL.Param1
		req.URL.RawQuery = urlKit.CombineQueryString(req.URL.RawQuery, urlKit.ToQueryString(extraQuery))
	}
	proxy := &httputil.ReverseProxy{
		Director: director,
		ErrorLog: errorLogger,
		ErrorHandler: func(rw http.ResponseWriter, req *http.Request, e error) {
			err = e
		},
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
	return err
}

func GetReqUrlPathForForward(req *http.Request, contextPath string) *string {
	contextPath = httpKit.PolyfillContextPath(contextPath)
	switch contextPath {
	case "":
		fallthrough
	case "/":
		return nil
	default:
		tmp := req.URL.Path
		if strKit.StartWith(tmp, contextPath) {
			tmp = strKit.SubAfter(tmp, len(contextPath))
			tmp = strKit.PrependIfMissing(tmp, "/")
			return strKit.GetStringPtr(tmp)
		}
		return nil
	}
}
