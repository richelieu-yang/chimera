package processKit

import "github.com/gogf/gf/v2/os/gproc"

var (
	// NewManager
	NewManager func() *gproc.Manager = gproc.NewManager

	// NewProcess
	NewProcess func(path string, args []string, environment ...[]string) *gproc.Process = gproc.NewProcess

	// NewProcessCmd
	NewProcessCmd func(cmd string, environment ...[]string) *gproc.Process = gproc.NewProcessCmd
)
