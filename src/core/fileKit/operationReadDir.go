package fileKit

import "os"

// ReadDir 获取指定目录下的文件或目录（另一种遍历目录的方法）.
/*
PS:
(1) 不包含子目录下的文件（或目录）.

@param name 目录路径
			(1) 如果为"" ，将返回error
			(2) 可以为"."
@return 第一个返回值是按照 文件名 从小到大 排序的
*/
var ReadDir func(name string) ([]os.DirEntry, error) = os.ReadDir
