package idKit

import (
	"encoding/base64"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/richelieu-yang/chimera/v2/src/atomicKit"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
	"strconv"
	"sync"
)

var once sync.Once
var i *gtype.Int
var prefix string

// SetLocalUniqueIdPrefix 设置本地唯一id的前缀（默认""）.
func SetLocalUniqueIdPrefix(p string) {
	prefix = p
}

// NewLocalUniqueId 生成本地唯一id（肯定不会重复）.
/*
	e.g.
		{
			id := idKit.NewLocalUniqueId()
			fmt.Println(id)                                                                                // MQ
			fmt.Println(base64Kit.DecodeStringToString(id, base64Kit.WithEncoding(base64.RawURLEncoding))) // 1 <nil>
		}
		{
			id := idKit.NewLocalUniqueId()
			fmt.Println(id)                                                                                // Mg
			fmt.Println(base64Kit.DecodeStringToString(id, base64Kit.WithEncoding(base64.RawURLEncoding))) // 2 <nil>
		}
*/
func NewLocalUniqueId() string {
	once.Do(func() {
		i = atomicKit.NewInt(0)
	})

	str := strconv.Itoa(i.Add(1))
	str = base64Kit.EncodeStringToString(str, base64Kit.WithEncoding(base64.RawURLEncoding))
	return prefix + str
}
