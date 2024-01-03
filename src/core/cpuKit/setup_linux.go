//go:build linux

package cpuKit

import (
	_ "go.uber.org/automaxprocs"
)

func SetUp() {
	// go.uber.org/automaxprocs作用: Linux环境下，自动设置 GOMAXPROCS 的值，以便更好地利用容器的CPU资源.
}
