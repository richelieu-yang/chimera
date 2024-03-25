package otelKit

import "github.com/richelieu-yang/chimera/v3/src/core/errorKit"

var (
	NotOtelRequestError = errorKit.Newf("not otel request")

	NotSetupError = errorKit.Newf("Haven’t been set up correctly")
)
