package otel

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	otlptracegrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var otelEndpoint = os.Getenv("OTEL_GRPC_ENDPOINT")

func SetupOTel(ctx context.Context) error {
	// Create an OTLP exporter (gRPC)
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpointURL(otelEndpoint), //otlptracegrpc.WithEndpoint("localhost:4317")
	)

	if err != nil {
		return fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	// Define resource attributes
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("my-service"),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	// Create a TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	// Set the global tracer provider
	otel.SetTracerProvider(tp)

	return nil
}
