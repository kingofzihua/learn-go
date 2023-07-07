package main

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

//nolint:deadcode,unused
func initOtlpProvider(ctx context.Context, cfg *Otlp, serviceName, environment string) (*tracesdk.TracerProvider, error) {
	if !cfg.Enabled {
		return nil, nil
	}

	if cfg.Endpoint == "" || cfg.Token == "" {
		return nil, nil
	}

	endpoint, token, samplerRatio := cfg.Endpoint, cfg.Token, float64(cfg.SamplerRatio)

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
		// 鉴权信息
		otlptracegrpc.WithHeaders(map[string]string{
			"Authentication": token,
		}),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))

	traceExp, err := otlptrace.New(ctx, traceClient)

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

	bsp := tracesdk.NewBatchSpanProcessor(traceExp)

	sampler := tracesdk.ParentBased(tracesdk.TraceIDRatioBased(samplerRatio))

	tracerProvider := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(sampler),
		tracesdk.WithSpanProcessor(bsp),
		tracesdk.WithResource(res),
	)

	return tracerProvider, err
}
