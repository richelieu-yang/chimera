package browserKit

import (
	"testing"
)

// 让默认浏览器打开 "https://www.baidu.com"
func TestOpenURL(t *testing.T) {
	if err := OpenURL("https://www.baidu.com"); err != nil {
		panic(err)
	}
}
