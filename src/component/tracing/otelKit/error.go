package otelKit

import "github.com/richelieu-yang/chimera/v2/src/core/errorKit"

var (
	NotOtelRequestError = errorKit.New("not otel request")
)
