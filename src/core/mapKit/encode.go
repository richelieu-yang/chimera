package mapKit

import (
	"github.com/fatih/structs"
)

// Encode 结构体 => map[string]interface{}
/*
@param obj 	(1) 结构体实例 || 结构体实例的指针
			(2) 不能为nil（否则会panic）
@return 必定不为nil
*/
func Encode(obj interface{}) map[string]interface{} {
	return structs.Map(obj)
}
