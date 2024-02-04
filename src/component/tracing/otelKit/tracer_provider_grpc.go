package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewGrpcTracerProvider
/*
PS:
(1) 使用 otlptracegrpc;
(2) 此方法第二个返回值为nil的情况下，建议立即调用 otel.SetTracerProvider.

@param endpoint 	可以为""，将采用默认值: "localhost:4317"
@param serviceName 	服务名
@param attributeMap	可以为nil
@param opts 		(1) 额外配置（不要涉及 otlptracegrpc.WithEndpoint，因为在此处配了也没用）
					(2) otlptracegrpc.WithInsecure(): 	(a) 配置的话，将使用非安全协议（http...）;
														(b) 不配置的话，将使用安全协议（https...）.
					(3) 一般情况下，传个 otlptracegrpc.WithInsecure() 就足够了.
*/
func NewGrpcTracerProvider(endpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) (*trace.TracerProvider, error) {
	// 放在最后面（优先级最高）
	if strKit.IsNotEmpty(endpoint) {
		opts = append(opts, otlptracegrpc.WithEndpoint(endpoint))
	}
	// 创建 exporter 实例
	exporter, err := otlptracegrpc.New(context.TODO(), opts...)
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
		trace.WithSampler(trace.AlwaysSample()),
	)
	return tp, nil
}
