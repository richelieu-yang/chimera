package regexpKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"regexp"
)

// StringToRegexp
/*
PS: 主要是处理 传参str 中的"*".

@param str 			(1) 不会进行trim操作; (2) 不同于原生的 regexp.Compile().
@param borderArgs 	边界匹配
*/
func StringToRegexp(str string, borderArgs ...bool) (*regexp.Regexp, error) {
	startBorder := false
	endBorder := false

	if borderArgs != nil {
		switch len(borderArgs) {
		case 1:
			startBorder = borderArgs[0]
		case 2:
			fallthrough
		default:
			startBorder = borderArgs[0]
			endBorder = borderArgs[1]
		}
	}

	/* 处理"." */
	str = strKit.ReplaceAll(str, ".", "\\.")

	/* 处理"*" */
	if strKit.Index(str, "*") != -1 {
		// 额外兼容情况：连续"*"的情况
		starRegexp := regexp.MustCompile("\\*+")
		str = starRegexp.ReplaceAllString(str, ".+")
	}
	if startBorder {
		str = strKit.PrependIfMissing(str, "^")
	}
	if endBorder {
		str = strKit.AppendIfMissing(str, "$")
	}

	return regexp.Compile(str)
}
