package unicodeKit

import (
	"strconv"
	"strings"
)

// Encode
/*
@param str 支持中文

e.g.
("测试") => "\u6d4b\u8bd5"
*/
func Encode(str string) string {
	textQuoted := strconv.QuoteToASCII(str)
	// 去掉首尾的双引号
	return textQuoted[1 : len(textQuoted)-1]
}

func Decode(str string) (rst string, err error) {
	return strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
}
