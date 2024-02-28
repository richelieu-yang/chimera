package i18nPackKit

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	SetUp(language.English, nil)

	/* 英文 */
	_, err := LoadMessageFile("_test.en.properties")
	if err != nil {
		panic(err)
	}
	/* 中文（简体） */
	{
		mf, err := LoadMessageFile("_test.zh-Hans.properties")
		if err != nil {
			panic(err)
		}
		if err := Associate(mf, "zh", "zh-CN"); err != nil {
			panic(err)
		}
	}
	/* 中文（繁体） */
	{
		mf, err := LoadMessageFile("_test.zh-Hant.properties")
		if err != nil {
			panic(err)
		}
		if err := Associate(mf, "zh-HK", "zh-TW", "zh-MO", "zh-SG"); err != nil {
			panic(err)
		}
	}

	fmt.Println(Seal(nil, "0", nil))               // {"code":"0","message":"No error."} <nil>
	fmt.Println(Seal([]string{"en"}, "0", nil))    // {"code":"0","message":"No error."} <nil>
	fmt.Println(Seal([]string{"zh"}, "0", nil))    // {"code":"0","message":"无错误。"} <nil>
	fmt.Println(Seal([]string{"zh-SG"}, "0", nil)) // {"code":"0","message":"無錯誤。"} <nil>
}
