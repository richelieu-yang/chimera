package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.File("")) // Key: '' Error:Field validation for '' failed on the 'file' tag

	// 目录存在
	fmt.Println(validateKit.File("chimera-lib"))                                         // Key: '' Error:Field validation for '' failed on the 'file' tag
	fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/chimera-lib")) // Key: '' Error:Field validation for '' failed on the 'file' tag
	// 文件存在
	fmt.Println(validateKit.File("chimera-lib/config.yaml"))                                         // <nil>
	fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml")) // <nil>
	// 文件不存在
	fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/chimera-lib/config111.yaml")) // Key: '' Error:Field validation for '' failed on the 'file' tag
	// 无效的文件路径
	fmt.Println(validateKit.File("chimera-lib\\config.yaml")) // Key: '' Error:Field validation for '' failed on the 'file' tag
}
