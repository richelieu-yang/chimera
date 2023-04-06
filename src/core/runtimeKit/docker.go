package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/docker"
)

// GetDockerIdList
/*
通过第三方库 gopsutil 实现.
*/
func GetDockerIdList() ([]string, error) {
	return docker.GetDockerIDList()
}
