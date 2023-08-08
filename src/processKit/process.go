package processKit

import (
	"github.com/shirou/gopsutil/v3/process"
	"os"
)

var PID int = os.Getpid()

var GetRunningPids func() ([]int32, error) = process.Pids

// PidExists pid是否存在?
var PidExists func(pid int32) (bool, error) = process.PidExists
