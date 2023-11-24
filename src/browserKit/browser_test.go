package browserKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"testing"
	"time"
)

// 让默认浏览器打开 "https://www.baidu.com"
func TestOpenURL(t *testing.T) {
	if err := OpenURL("https://www.baidu.com"); err != nil {
		panic(err)
	}
}

func TestOpenFile(t *testing.T) {
	temp, err := pathKit.GetExclusiveTempDir()
	if err != nil {
		panic(err)
	}
	path := pathKit.Join(temp, fmt.Sprintf("%s.txt", idKit.NewXid()))

	f, err := fileKit.Create(path)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(`份额额服务强
无敌
aas
~!@#$%^&*()_+{}:"|<>?"
群无
多`)
	if err != nil {
		panic(err)
	}
	_ = f.Close()

	if err := OpenFile(path); err != nil {
		panic(err)
	}

	// Richelieu: 过一会再删，防止打开后内容为空
	time.Sleep(time.Second * 3)
	_ = fileKit.Delete(path)
}
