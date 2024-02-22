package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/config/propertiesKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func main() {
	data, err := fileKit.ReadFile("demo.properties")
	if err != nil {
		panic(err)
	}

	var raw interface{}
	if err := propertiesKit.Unmarshal(data, &raw); err != nil {
		panic(err)
	}
	fmt.Println(raw)

	//bundle := i18n.NewBundle(language.English)
	//bundle.RegisterUnmarshalFunc("json", jsoniter.Unmarshal)
	//_, err := bundle.LoadMessageFile("active.en.json")
	//if err != nil {
	//	panic(err)
	//}
	//_, err = bundle.LoadMessageFile("active.zh-CN.json")
	//if err != nil {
	//	panic(err)
	//}
	//
	////localizer := i18n.NewLocalizer(bundle, "en")
	////localizer := i18n.NewLocalizer(bundle, "zh-CN")
	//localizer := i18n.NewLocalizer(bundle, "111")
	//helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
	//	DefaultMessage: &i18n.Message{
	//		ID: "HelloPerson",
	//		//Other: "Hello {{.Name}}",
	//	},
	//	//TemplateData: map[string]interface{}{
	//	//	"Name": "Nick",
	//	//},
	//})
	//fmt.Println(helloPerson) // 输出: Hello Nick
}
