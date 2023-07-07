package main

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initJaegerProvider(ctx context.Context, cfg *Jaeger, serviceName string, environment string) (*tracesdk.TracerProvider, error) {
	if !cfg.Enabled {
		return nil, nil
	}

	if cfg.Host == "" || cfg.Port == "" {
		return nil, nil
	}

	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost(cfg.Host),
		jaeger.WithAgentPort(cfg.Port),
	))

	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", environment),
		),
	)

	if err != nil {
		return nil, err
	}

	samplerRatio := float64(cfg.SamplerRatio)

	sampler := tracesdk.ParentBased(tracesdk.TraceIDRatioBased(samplerRatio))

	provider := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(sampler),
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(res),
	)

	return provider, nil
}
