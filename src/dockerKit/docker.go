package dockerKit

import (
	"github.com/shirou/gopsutil/v3/docker"
)

// GetDockerIdList returns a list of DockerID.
var GetDockerIdList func() ([]string, error) = docker.GetDockerIDList
