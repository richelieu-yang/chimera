package otelKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	"time"
)

var tp *trace.TracerProvider

func MustSetUp(grpcEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) {
	err := SetUp(grpcEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(grpcEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) (err error) {
	if err = strKit.AssertNotEmpty(serviceName, "serviceName"); err != nil {
		return
	}

	tp, err = NewGrpcTracerProvider(grpcEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		return
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tp)

	logrus.RegisterExitHandler(func() {
		ShutdownTracerProvider(tp, time.Second*3)
	})
	return
}

func check() error {
	if tp == nil {
		return NotSetupError
	}
	return nil
}
