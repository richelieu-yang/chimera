package mapKit

import (
	"gitee.com/richelieu042/go-scales/src/core/pointerKit"
	"github.com/mitchellh/mapstructure"
)

// DecodeToPointer 结构体（包括map）的实例或指针 => 结构体（包括map）的指针
/*
PS:
(1) 也可用于: map转map；
(2) 字段标签(tag)不同于json，详见: https://mp.weixin.qq.com/s/n_AXYv-p_ev-q23mhpDkOg .

@param input	structure（包括map）
@param output 	指针，must be a pointer to a map or struct
*/
func DecodeToPointer(input interface{}, outputPtr interface{}) error {
	err := pointerKit.AssertPointer(outputPtr, "outputPtr")
	if err != nil {
		return err
	}
	return mapstructure.Decode(input, outputPtr)
}
