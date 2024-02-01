package otelKit

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewHttpTracerProvider 通过otlptracehttp
/*
PS: 此方法第二个返回值为nil的情况下，建议立即调用 otel.SetTracerProvider.

@param endpoint 	e.g."localhost:4318"
@param serviceName 	服务名
@param attributeMap	可以为nil？
@param opts 		额外配置（不建议涉及 otlptracehttp.WithEndpoint，因为在此处配了也没用）
*/
func NewHttpTracerProvider(endpoint, serviceName string, attributeMap map[string]string, opts ...otlptracehttp.Option) (*trace.TracerProvider, error) {
	// 默认: 使用非安全协议（http）
	opts = append([]otlptracehttp.Option{otlptracehttp.WithInsecure()}, opts...)
	// 放在最后面（优先级最高）
	opts = append(opts, otlptracehttp.WithEndpoint(endpoint))
	exporter, err := otlptracehttp.New(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	res := NewResourceWithAttributes(serviceName, attributeMap)

	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		trace.WithResource(res),
	)
	return tp, nil
}
