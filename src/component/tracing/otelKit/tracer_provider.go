package otelKit

import (
	"context"
	"github.com/sirupsen/logrus"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"

	"time"
)

// NewNoopTracerProvider returns a TracerProvider that does not record any telemetry.
func NewNoopTracerProvider() trace.TracerProvider {
	//return trace.NewNoopTracerProvider()
	return noop.NewTracerProvider()
}

func ShutdownTracerProvider(tp *sdktrace.TracerProvider, timeout time.Duration) {
	if tp == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Fail to shutdown tracer provider.")
	}
}
