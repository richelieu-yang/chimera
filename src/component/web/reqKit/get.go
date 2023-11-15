package reqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
)

func Get(url string, queryParams map[string][]string) (code int, data []byte, err error) {
	url, err = urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return
	}

	client := GetDefaultClient()
	resp := client.Get(url).Do()
	err = resp.Err
	if err != nil {
		return
	}

	// 不需要手动关闭
	//defer resp.Body.Close()

	code = resp.StatusCode
	data = resp.Bytes()
	if !resp.IsSuccessState() {
		err = errorKit.New("not success state(%d)", code)
	}
	return
}
