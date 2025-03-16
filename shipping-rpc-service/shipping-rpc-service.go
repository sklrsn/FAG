package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/sklrsn/gRPC-defs/shipping"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

type ShippingStatus string

func (os *ShippingStatus) String() string {
	return string(*os)
}

var (
	SUCCESS = ShippingStatus("SUCCESS")
)

type ShippingRPCService struct {
	shipping.ShippingServer
}

var ServiceName = "shipping-rpc-service"

var (
	tracer = otel.Tracer(ServiceName)
)

func (ps ShippingRPCService) Dispatch(ctx context.Context, dr *shipping.DispatchRequest) (*shipping.DispatchResponse, error) {
	ctx, span := tracer.Start(ctx, "Dispatch")
	defer span.End()

	processDispatchRequest(ctx, dr) //do nothing

	return &shipping.DispatchResponse{
		ShippingId: uuid.NewString(),
	}, nil
}

func processDispatchRequest(ctx context.Context, dr *shipping.DispatchRequest) {}
func (ps ShippingRPCService) Hold(ctx context.Context, hr *shipping.HoldRequest) (*shipping.HoldResponse, error) {
	ctx, span := tracer.Start(ctx, "Hold")
	defer span.End()

	processHoldRequest(ctx, hr) //do nothing

	return &shipping.HoldResponse{}, nil
}

func processHoldRequest(ctx context.Context, dr *shipping.HoldRequest) {}

func (ps ShippingRPCService) Retract(ctx context.Context, rr *shipping.RetractRequest) (*shipping.RetractResponse, error) {
	ctx, span := tracer.Start(ctx, "Retract")
	defer span.End()

	processRetractRequest(ctx, rr)

	return &shipping.RetractResponse{}, nil
}

func processRetractRequest(ctx context.Context, dr *shipping.RetractRequest) {}

func main() {
	if err := SetupOTel(context.Background()); err != nil {
		log.Fatalf("%v", err)
	}

	listener, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("%v", err)
	}

	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, ShippingRPCService{})

	log.Fatalf("%v", grpcServer.Serve(listener))
}

var (
	otelEndpoint = os.Getenv("OTEL_ENDPOINT")
)

func SetupOTel(ctx context.Context) error {
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otelEndpoint), //otlptracegrpc.WithEndpoint("localhost:4317")
	)
	if err != nil {
		return fmt.Errorf("failed to create OTLP exporter: %w", err)
	}
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("shipping-rpc-service"),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	return nil
}
