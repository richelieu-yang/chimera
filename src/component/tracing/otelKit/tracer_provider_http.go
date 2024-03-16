package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewHttpTracerProvider
/*
Deprecated: Use NewGrpcTracerProvider instead.

PS:
(1) 使用 otlptracehttp;
(2) 此方法第二个返回值为nil的情况下，建议立即调用 otel.SetTracerProvider.

@param endpoint 	可以为""，将采用默认值: "localhost:4318"
@param serviceName 	服务名
@param attributeMap	可以为nil
@param opts 		额外配置（不要涉及 otlptracehttp.WithEndpoint，因为在此处配了也没用）
@param opts 		(1) 额外配置（不要涉及 otlptracegrpc.WithEndpoint，因为在此处配了也没用）
					(2) otlptracegrpc.WithInsecure(): 	(a) 配置的话，将使用非安全协议（http...）;
														(b) 不配置的话，将使用安全协议（https...）.
					(3) 一般情况下，传个 otlptracegrpc.WithInsecure() 就足够了.
*/
func NewHttpTracerProvider(endpoint, serviceName string, attributeMap map[string]string, opts ...otlptracehttp.Option) (*trace.TracerProvider, error) {
	if err := validateKit.Var(endpoint, "omitempty,hostname_port"); err != nil {
		return nil, errorKit.New("invalid http endpoint(%s)", endpoint)
	}
	// 放在最后面（优先级最高）
	if strKit.IsNotEmpty(endpoint) {
		opts = append(opts, otlptracehttp.WithEndpoint(endpoint))
	}
	// 创建 exporter 实例
	exporter, err := otlptracehttp.New(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	res, err := newDetailedResource(serviceName, attributeMap)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		trace.WithResource(res),
	)
	return tp, nil
}
