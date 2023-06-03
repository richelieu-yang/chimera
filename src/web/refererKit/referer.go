package refererKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/regexpKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"regexp"
)

type RefererVerifierBuilder struct {
	None    bool
	Blocked bool

	// Route 路由的正则字符串
	Route string
	// ServerNames referer白名单的正则字符串s
	ServerNames []string
}

func (builder *RefererVerifierBuilder) Build() (*RefererVerifier, error) {
	var err error

	v := &RefererVerifier{}
	v.routeRegexp, err = regexpKit.StringToRegexp(builder.Route)
	if err != nil {
		return nil, err
	}

	/* refererRegexps 属性 */
	serverNames := builder.ServerNames
	// 优化 serverNames
	serverNames = sliceKit.RemoveEmpty(serverNames, true)
	serverNames = sliceKit.Uniq(serverNames)
	refererRegexps := make([]*regexp.Regexp, 0, len(serverNames))
	for _, serverName := range serverNames {
		var tmp *regexp.Regexp
		tmp, err = regexpKit.StringToRegexp(serverName)
		if err != nil {
			return nil, err
		}
		refererRegexps = append(refererRegexps, tmp)
	}
	v.refererRegexps = refererRegexps

	v.none = builder.None
	v.blocked = builder.Blocked
	return v, nil
}

type RefererVerifier struct {
	/*
		必定不为nil
	*/
	routeRegexp    *regexp.Regexp
	none           bool
	blocked        bool
	refererRegexps []*regexp.Regexp
}

// VerifyByGinContext 验证referer
func (verifier *RefererVerifier) VerifyByGinContext(ctx *gin.Context) (bool, string) {
	route := ctx.Request.URL.Path
	referer := ctx.GetHeader("Referer")

	return verifier.Verify(route, referer)
}

// Verify 验证referer
/*
@param route	请求的路由（可以通过 ctx.FullPath() 获取）
@param referer 	请求的referer（可能为""）
@return 验证是否通过 + 验证失败的原因
*/
func (verifier *RefererVerifier) Verify(route, referer string) (bool, string) {
	if verifier == nil {
		return false, "verifier == nil"
	}

	/* route */
	if !verifier.routeRegexp.MatchString(route) {
		// 路由不匹配的情况下，默认通过referer验证
		return true, ""
	}

	/* referer */
	if strKit.IsEmpty(referer) {
		return verifier.none, "none"
	}

	var prefix string
	if strKit.StartWith(referer, "http://") {
		prefix = "http://"
	} else if strKit.StartWith(referer, "https://") {
		prefix = "https://"
	} else {
		return verifier.blocked, "blocked"
	}

	referer = strKit.RemovePrefixIfExist(referer, prefix)
	referer = strKit.SubBeforeString(referer, "/")
	// 忽略端口号（有的话）
	referer = strKit.SubBeforeString(referer, ":")

	for _, re := range verifier.refererRegexps {
		if re.MatchString(referer) {
			return true, ""
		}
	}
	return false, "referer no matched"
}
