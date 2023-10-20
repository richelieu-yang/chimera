package otelKit

import "go.opentelemetry.io/otel/metric"

var NewMeterConfig func(opts ...metric.MeterOption) metric.MeterConfig = metric.NewMeterConfig

var NewAddConfig func(opts []metric.AddOption) metric.AddConfig = metric.NewAddConfig
