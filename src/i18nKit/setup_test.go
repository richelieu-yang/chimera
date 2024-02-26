package i18nKit

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	MustSetUp(language.English, "_test.en.properties", "_test.zh.properties")

	fmt.Println(GetMessage("0"))       // no error <nil>
	fmt.Println(GetMessage("0", ""))   // no error <nil>
	fmt.Println(GetMessage("0", "zh")) // 无错误 <nil>
	fmt.Println(GetMessage("0", "en")) // no error <nil>
}
