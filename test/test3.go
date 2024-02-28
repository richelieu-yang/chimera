package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/langKit"
)

func main() {
	fmt.Println(langKit.S2T("鼠标"))
	fmt.Println(langKit.S2HK("鼠标"))
	fmt.Println(langKit.S2TW("鼠标"))

	//s2t, err := gocc.New("s2t")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//in := `自然语言处理是人工智能领域中的一个重要方向。`
	//out, err := s2t.Convert(in)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n%s\n", in, out)
	////自然语言处理是人工智能领域中的一个重要方向。
	////自然語言處理是人工智能領域中的一個重要方向。
}
