package i18nPackKit

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	MustSetUp(language.Chinese, []string{
		"_test.en.properties",
		"_test.zh.properties",
	}, nil)

	fmt.Println(Seal(nil, "0", nil))            // {"code":"0","message":"无错误"} <nil>
	fmt.Println(Seal([]string{"en"}, "0", nil)) // {"code":"0","message":"no error"} <nil>
	fmt.Println(Seal([]string{"zh"}, "0", nil)) // {"code":"0","message":"无错误"} <nil>
}
