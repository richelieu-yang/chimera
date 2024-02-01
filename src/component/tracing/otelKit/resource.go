package otelKit

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.22.0"
)

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
