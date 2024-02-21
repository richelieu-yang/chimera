package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// 初始化翻译包，通常会从文件加载资源
	bundle := i18n.NewBundle(language.English)
	bundle.MustParseMessageFileBytes([]byte(`hello = Hello, world!`), "en-US")
	bundle.MustParseMessageFileBytes([]byte(`hello = Bonjour le monde!`), "fr-FR")

	// 创建一个本地化器（localizer），用于获取指定语言的翻译
	localizer := i18n.NewLocalizer(bundle, "")

	// 设置用户偏好语言
	//tag, _ := language.Parse("fr-FR") // 假设用户选择法语

	// 获取翻译并打印
	msg := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "hello",
	})
	fmt.Println(msg) // 输出：Bonjour le monde!
}
