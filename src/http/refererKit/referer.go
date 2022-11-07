package refererKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"regexp"
)

type (
	RefererVerifier struct {
		/*
			!= nil: 正则匹配的才要验证referer；正则匹配的直接认为验证referer通过
			== nil: 不判断route，直接验证referer
		*/
		routeRegexp    *regexp.Regexp
		none           bool
		blocked        bool
		refererRegexps []*regexp.Regexp
	}
)

func NewRefererVerifier(route string, serverNames []string, none bool, blocked bool) (*RefererVerifier, error) {
	instance := &RefererVerifier{
		routeRegexp:    nil,
		none:           true,
		blocked:        false,
		refererRegexps: nil,
	}

	route = strKit.Trim(route)
	if strKit.IsNotEmpty(route) {

	} else {
		instance.routeRegexp = nil
	}

	// TODO:
	return nil, nil
}

// Verify 验证referer
/*
@param route	请求的路由（可以通过 ctx.FullPath() 获取）
@param referer 	请求的referer（可能为""）
@return 验证是否通过 + 验证失败的原因
*/
func (v *RefererVerifier) Verify(route, referer string) (bool, string) {
	if v == nil {
		return false, "v == nil"
	}

	/* route */
	if v.routeRegexp != nil {
		if v.routeRegexp.MatchString(route) {
			// 对于当前路由，需要验证referer，继续向下执行
		} else {
			// 对于当前路由，无需验证referer
			return true, ""
		}
	} else {
		// 对于当前路由，需要验证referer，继续向下执行
	}

	/* referer */
	if strKit.IsEmpty(referer) {
		return v.none, "none"
	}

	var prefix string
	if strKit.StartWith(referer, "http://") {
		prefix = "http://"
	} else if strKit.StartWith(referer, "https://") {
		prefix = "https://"
	} else {
		return v.blocked, "blocked"
	}

	referer = strKit.RemovePrefixIfExist(referer, prefix)
	referer = strKit.SubBeforeString(referer, "/")
	// 忽略端口号（有的话）
	referer = strKit.SubBeforeString(referer, ":")

	//if sliceKit.ContainsStringIgnoreCase(v.serverNames, referer) {
	//	return true, ""
	//}
	for _, re := range v.refererRegexps {
		if re.MatchString(referer) {
			return true, ""
		}
	}
	return false, "referer no matched"
}
