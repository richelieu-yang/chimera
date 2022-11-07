package refererKit

import (
	"github.com/richelieu42/go-scales/src/core/regexpKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"regexp"
)

type (
	RefererVerifier struct {
		/*
			必定不为nil
		*/
		routeRegexp    *regexp.Regexp
		none           bool
		blocked        bool
		refererRegexps []*regexp.Regexp
	}
)

func NewRefererVerifier(route string, serverNames []string, none bool, blocked bool) (v *RefererVerifier, err error) {
	v = &RefererVerifier{}

	v.routeRegexp, err = regexpKit.StringToRegexp(route)
	if err != nil {
		return
	}

	refererRegexps := make([]*regexp.Regexp, len(serverNames))
	for _, serverName := range serverNames {
		var tmp *regexp.Regexp
		tmp, err = regexpKit.StringToRegexp(serverName)
		if err != nil {
			return
		}
		refererRegexps = append(refererRegexps, tmp)
	}
	v.refererRegexps = refererRegexps

	v.none = none
	v.blocked = blocked
	return
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
	if !v.routeRegexp.MatchString(route) {
		// 路由不匹配的情况下，默认通过referer验证
		return true, ""
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
