package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"regexp"
)

func main() {
	str := "https://blog.csdn.net/weixin_44014995/article/details/120332422"

	re := regexp.MustCompile("\\\\")
	fmt.Println(re.Match([]byte("\\"))) // true
}

// VerifyReferer 验证 referer
/*
@return 验证是否成功 + 验证失败的原因
*/
func VerifyReferer(referer string, none bool, blocked bool) (bool, string) {
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

	i := strKit.Index(referer, "/")
	if i != -1 {
		referer = strKit.SubBefore(referer, i)
	}
	i = strKit.Index(referer, ":")
	if i != -1 {
		referer = strKit.SubBefore(referer, i)
	}

}
