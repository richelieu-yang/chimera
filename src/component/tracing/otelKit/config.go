package otelKit

type JaegerConfig struct {
	// GrpcEndpoint 可以为""（采用默认值"localhost:4317"）
	GrpcEndpoint string `json:"grpcEndpoint" yaml:"grpcEndpoint"`
}
