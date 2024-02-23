package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", jsoniter.Unmarshal)
	_, err := bundle.LoadMessageFile("active.en.json")
	if err != nil {
		panic(err)
	}
	_, err = bundle.LoadMessageFile("active.zh-CN.json")
	if err != nil {
		panic(err)
	}

	//localizer := i18n.NewLocalizer(bundle, "en")
	//localizer := i18n.NewLocalizer(bundle, "zh-CN")
	localizer := i18n.NewLocalizer(bundle, "111")
	helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "HelloPerson",
			//Other: "Hello {{.Name}}",
		},
		//TemplateData: map[string]interface{}{
		//	"Name": "Nick",
		//},
	})
	fmt.Println(helloPerson) // 输出: Hello Nick
}
