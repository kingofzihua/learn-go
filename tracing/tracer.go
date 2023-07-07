package main

import (
	"context"
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewTracer return trace.TracerProvider and cleanup
func NewTracer(jaeger *Jaeger) (*tracesdk.TracerProvider, func(), error) {
	//tracerProvider, err := initOtlpProvider(ctx, cfg.Otlp, cfg.GetAppName(), cfg.GetAppEnv())
	ctx := context.Background()

	tracerProvider, err := initJaegerProvider(ctx, jaeger, "app_name", "prod")

	if err != nil {
		return nil, nil, err
	}

	// set global tracer provider
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, func() {
		_ = tracerProvider.Shutdown(context.Background())
	}, nil
}
