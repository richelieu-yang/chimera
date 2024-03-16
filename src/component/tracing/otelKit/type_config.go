package otelKit

type JaegerConfig struct {
	// Type "grpc" || "http"
	Type string `json:"type" yaml:"type"`

	// Endpoint 可以为""，将采用默认值 "localhost:4317"(grpc) || "localhost:4318"（http）
	Endpoint string `json:"grpcEndpoint" yaml:"grpcEndpoint"`
}
