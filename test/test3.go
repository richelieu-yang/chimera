package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"go.opentelemetry.io/otel/trace/noop"
)

func main() {
	tp := noop.NewTracerProvider()
}
