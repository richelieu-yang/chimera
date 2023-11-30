package envKit

import (
	"github.com/joho/godotenv"
	"io"
)

var (
	// Load
	/*
		PS:
		(1) 默认情况下，加载的是项目根目录下的.env文件;
		(2) 如果多个文件中存在同一个键，那么先出现的优先，后出现的不生效;
		(3) 会存储到程序的环境变量中.
	*/
	Load func(filenames ...string) (err error) = godotenv.Load

	// Read
	/*

	 */
	Read func(filenames ...string) (envMap map[string]string, err error) = godotenv.Read

	// Unmarshal
	/*

	 */
	Unmarshal func(str string) (envMap map[string]string, err error) = godotenv.Unmarshal

	// Parse
	/*

	 */
	Parse func(r io.Reader) (map[string]string, error) = godotenv.Parse
)
