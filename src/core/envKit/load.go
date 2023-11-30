package envKit

import (
	"github.com/joho/godotenv"
	"io"
)

var (
	// Load 从 文件 中读取配置，并存储到程序的环境变量中.
	/*
		PS:
		(1) 默认情况下，加载的是项目根目录下的.env文件;
		(2) 如果多个文件中存在同一个键，那么先出现的优先，后出现的不生效;
		(3) 会存储到程序的环境变量中.
	*/
	Load func(filenames ...string) (err error) = godotenv.Load

	// Read 从 文件 中读取配置.
	/*

	 */
	Read func(filenames ...string) (envMap map[string]string, err error) = godotenv.Read

	// Unmarshal 从 string 中读取配置.
	/*

	 */
	Unmarshal func(str string) (envMap map[string]string, err error) = godotenv.Unmarshal

	// Parse 从 io.Reader 中读取配置.
	/*

	 */
	Parse func(r io.Reader) (map[string]string, error) = godotenv.Parse
)

func Load(filenames ...string) (err error) {
	filenames = filenamesOrDefault(filenames)

	for _, filename := range filenames {
		err = loadFile(filename, false)
		if err != nil {
			return // return early on a spazout
		}
	}
	return
}
