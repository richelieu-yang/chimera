package otelKit

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/sdk/trace"
	"time"
)

func ShutdownTracerProvider(tp *trace.TracerProvider, timeout time.Duration) {
	if tp == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Fail to shutdown tracer provider.")
	}
}
