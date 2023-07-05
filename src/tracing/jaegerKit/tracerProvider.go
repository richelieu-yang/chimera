package jaegerKit

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// NewTracerProvider
/*
	参考: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go

	tracerProvider returns an OpenTelemetry TracerProvider configured to use
	the Jaeger exporter that will send spans to the provided url. The returned
	TracerProvider will also use a Resource configured with all the information
	about the application.

	@param url e.g."http://localhost:14268/api/traces"
*/
func NewTracerProvider(url string, rc *ResourceConfig) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(rc.ServiceName),
			attribute.String("environment", rc.Environment),
			attribute.Int64("ID", rc.ID),
		)),
	)
	return tp, nil
}
