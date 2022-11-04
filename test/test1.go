package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"regexp"
)

func main() {
	engine := gin.Default()

	engine.Any("/test.act", func(ctx *gin.Context) {
		ctx.String(200, timeKit.FormatCurrentTime())
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}

// VerifyReferer 验证 referer
/*
@return 验证是否通过 + 验证失败的原因
*/
func VerifyReferer(referer string, none bool, blocked bool, serverNames []string, regexps []*regexp.Regexp) (bool, string) {
	if strKit.IsEmpty(referer) {
		return none, "none"
	}

	var prefix string
	if strKit.StartWith(referer, "http://") {
		prefix = "http://"
	} else if strKit.StartWith(referer, "https://") {
		prefix = "https://"
	} else {
		return blocked, "blocked"
	}

	referer = strKit.RemovePrefixIfExist(referer, prefix)
	referer = strKit.SubBeforeString(referer, "/")
	// 忽略端口号（有的话）
	referer = strKit.SubBeforeString(referer, ":")

	if sliceKit.ContainsStringIgnoreCase(serverNames, referer) {
		return true, ""
	}
	for _, re := range regexps {
		if re.MatchString(referer) {
			return true, ""
		}
	}
	return false, "no match found"
}
