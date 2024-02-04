package otelKit

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.22.0"
)

func newDetailedResource(serviceName string, attributeMap map[string]string) (*resource.Resource, error) {
	attributes := []attribute.KeyValue{
		semconv.ServiceNameKey.String(serviceName),
	}
	for key, value := range attributeMap {
		attributes = append(attributes, attribute.String(key, value))
	}

	res, err := resource.New(context.TODO(),
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(attributes...),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// newResourceWithAttributes Record information about this application in a Resource.
func newResourceWithAttributes(serviceName string, attributeMap map[string]string) *resource.Resource {
	var attrs = []attribute.KeyValue{
		semconv.ServiceNameKey.String(serviceName),
	}
	for key, value := range attributeMap {
		attrs = append(attrs, attribute.String(key, value))
	}

	return resource.NewWithAttributes(semconv.SchemaURL, attrs...)
}
