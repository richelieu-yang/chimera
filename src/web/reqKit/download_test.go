package reqKit

import "testing"

func TestDownloadToFile(t *testing.T) {
	url := "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"

	if err := DownloadToFile(url, "aaa.png"); err != nil {
		panic(err)
	}
}
