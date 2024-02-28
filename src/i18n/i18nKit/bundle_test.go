package i18nKit

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestNewBundle(t *testing.T) {
	bundle := NewBundle(language.English)

	/* 英文 */
	_, err := bundle.LoadMessageFile("_test.en.properties")
	if err != nil {
		panic(err)
	}
	/* 中文（简体） */
	{
		mf, err := bundle.LoadMessageFile("_test.zh-Hans.properties")
		if err != nil {
			panic(err)
		}
		if err := Associate(bundle, mf, "zh", "zh-CN"); err != nil {
			panic(err)
		}
	}
	/* 中文（繁体） */
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
	fmt.Println("---")

	fmt.Println(GetMessage(bundle, "0", "zh-HK")) // 無錯誤。 <nil>
	fmt.Println(GetMessage(bundle, "0", "zh-TW")) // 無錯誤。 <nil>
	fmt.Println(GetMessage(bundle, "0", "zh-MO")) // 無錯誤。 <nil>
	fmt.Println(GetMessage(bundle, "0", "zh-SG")) // 無錯誤。 <nil>
	fmt.Println("---")

	/* 大小写不敏感 */
	fmt.Println(GetMessage(bundle, "0", "zh-hk")) // 無錯誤。 <nil>
	fmt.Println(GetMessage(bundle, "0", "zh-tW")) // 無錯誤。 <nil>
	fmt.Println("---")

	fmt.Println(GetMessage(bundle, "0", "en"))      // No error. <nil>
	fmt.Println(GetMessage(bundle, "0", "zh"))      // 无错误。 <nil>
	fmt.Println(GetMessage(bundle, "0", "zh-Hans")) // 无错误。 <nil>
	fmt.Println("---")

	/*
		第1个lang是无效的;
		第2个lang是无效的;
		第3个lang是有效的;
		最终采用了第3个lang.
	*/
	fmt.Println(GetMessage(bundle, "0", "", "zh-CCC", "en")) // No error. <nil>
}
