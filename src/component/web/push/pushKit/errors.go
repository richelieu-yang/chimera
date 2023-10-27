package pushKit

import "github.com/richelieu-yang/chimera/v2/src/core/errorKit"

var (
	ChannelClosedError = errorKit.New("Channel is already closed")
)
