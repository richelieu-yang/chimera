//go:build (linux || windows || darwin) && amd64

package jsonKit

import (
	"github.com/bytedance/sonic"
)

func init() {
	api = sonic.ConfigDefault
}
