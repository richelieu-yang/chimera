// Package ginKit
// 静态资源相关
package ginKit

import (
	"github.com/gin-gonic/gin"
)

// LoadHtmlFiles 加载（多个）html文件
func LoadHtmlFiles(engine IEngine, filePaths ...string) {
	engine.LoadHTMLFiles(filePaths...)
}

func LoadHtmlGlob(engine IEngine, pattern string) {
	engine.LoadHTMLGlob(pattern)
}

// StaticFile 静态资源（单个文件）
func StaticFile(group IGroup, relativePath, filePath string) {
	group.StaticFile(relativePath, filePath)
}

// StaticDir 静态资源（目录）
/*
@param relativePath	路由
@param dirPath		静态资源所在的目录（相对路径 || 绝对路径）
*/
func StaticDir(group IGroup, relativePath, dirPath string, listDirectory bool) {
	group.StaticFS(relativePath, gin.Dir(dirPath, listDirectory))
}
