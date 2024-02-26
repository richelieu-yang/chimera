package i18nKit

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestNewBundle(t *testing.T) {
	bundle, err := NewBundle(language.English, "_test.en.properties", "_test.zh.properties")
	if err != nil {
		panic(err)
	}

	fmt.Println(GetMessage(bundle, "0"))       // no error <nil>
	fmt.Println(GetMessage(bundle, "0", ""))   // no error <nil>
	fmt.Println(GetMessage(bundle, "0", "zh")) // 无错误 <nil>
	fmt.Println(GetMessage(bundle, "0", "en")) // no error <nil>

	/*
		第1个lang是无效的;
		第2个lang是有效的;
		最终采用了第2个lang.
	*/
	fmt.Println(GetMessage(bundle, "1", "zh-CN", "zh")) // 参数错误 <nil>
}
