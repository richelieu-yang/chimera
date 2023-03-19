package mapKit

import (
	"github.com/mohae/deepcopy"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

// Clone 浅拷贝.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	dolly := make(map[K]V)
	for k, v := range m {
		dolly[k] = v
	}
	return dolly
}

// DeepClone 深拷贝.
/*
PS:
(1) 参考: 「Go工具箱」推荐一个非常简单的深拷贝工具：deepcopy https://mp.weixin.qq.com/s/e3bL1i6WT-4MwK-SEpUa6Q；
(2) 不要使用 copier 来拷贝map，因为不管如何配置都是浅拷贝（感觉是bug）.
*/
func DeepClone[K comparable, V any](m map[K]V) (map[K]V, error) {
	obj := deepcopy.Copy(m)

	if dolly, ok := obj.(map[K]V); ok {
		return dolly, nil
	}
	return nil, errorKit.Simple("fail to deep clone")
}
