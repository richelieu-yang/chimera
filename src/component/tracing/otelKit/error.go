package otelKit

import "github.com/richelieu-yang/chimera/v3/src/core/errorKit"

var (
	NotOtelRequestError = errorKit.New("not otel request")

	NotSetupError = errorKit.New("Haven’t been set up correctly")
)
