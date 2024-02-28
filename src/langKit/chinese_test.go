package langKit

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(T2S("裏 鼠標 臺灣 香港"))

	fmt.Println(S2T("里 鼠标 台湾 香港"))
	fmt.Println(S2HK("里 鼠标 台湾 香港"))
	fmt.Println(S2TW("里 鼠标 台湾 香港"))
}
