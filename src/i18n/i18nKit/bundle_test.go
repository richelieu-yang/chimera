package i18nKit

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestNewBundle(t *testing.T) {
	bundle := NewBundle(language.English)

	_, err := bundle.LoadMessageFile("_test.en.properties")
	if err != nil {
		panic(err)
	}
	{
		mf, err := bundle.LoadMessageFile("_test.zh-Hans.properties")
		if err != nil {
			panic(err)
		}
		if err := Associate(bundle, mf, "zh", "zh-CN"); err != nil {
			panic(err)
		}
	}
	{
		mf, err := bundle.LoadMessageFile("_test.zh-Hant.properties")
		if err != nil {
			panic(err)
		}
		if err := Associate(bundle, mf, "zh-HK", "zh-TW", "zh-MO", "zh-SG"); err != nil {
			panic(err)
		}
	}

	fmt.Println(GetMessage(bundle, "0"))     // No error. <nil>
	fmt.Println(GetMessage(bundle, "0", "")) // No error. <nil>

	fmt.Println(GetMessage(bundle, "0", "zh-TW"))   // 無錯誤。 <nil>
	fmt.Println(GetMessage(bundle, "0", "en"))      // No error. <nil>
	fmt.Println(GetMessage(bundle, "0", "zh"))      // 无错误。 <nil>
	fmt.Println(GetMessage(bundle, "0", "zh-Hans")) // 无错误。 <nil>

	/*
		第1个lang是无效的;
		第2个lang是无效的;
		第3个lang是有效的;
		最终采用了第3个lang.
	*/
	fmt.Println(GetMessage(bundle, "0", "", "zh-CCC", "en")) // No error. <nil>
}
