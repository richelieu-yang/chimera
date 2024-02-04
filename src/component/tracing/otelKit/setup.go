package otelKit

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"time"
)

func MustSetUp(grpcEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) {
	err := SetUp(grpcEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(grpcEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) error {
	tp, err := NewGrpcTracerProvider(grpcEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		return err
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tp)

	logrus.RegisterExitHandler(func() {
		ShutdownTracerProvider(tp, time.Second*3)
	})

	return nil
}
