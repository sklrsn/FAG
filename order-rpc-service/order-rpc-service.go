package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/sklrsn/gRPC-defs/order"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

type OrderStatus string

func (os *OrderStatus) String() string {
	return string(*os)
}

var (
	SUCCESS = OrderStatus("SUCCESS")
)

type OrderRPCService struct {
	order.OrderEngineServer
}

var ServiceName = "order-rpc-service"

var (
	tracer = otel.Tracer(ServiceName)
)

func (os OrderRPCService) Reserve(ctx context.Context, rr *order.ReserveRequest) (*order.ReserveResponse, error) {
	ctx, span := tracer.Start(ctx, "Reserve")
	defer span.End()

	processReserverRequest(ctx, rr) // do nothing

	return &order.ReserveResponse{
		OrderReservationId: uuid.NewString(),
	}, nil
}

func processReserverRequest(ctx context.Context, rr *order.ReserveRequest) {}

func (os OrderRPCService) Release(ctx context.Context, rr *order.ReleaseRequest) (*order.ReleaseResponse, error) {
	ctx, span := tracer.Start(ctx, "Reserve")
	defer span.End()

	processReleaseRequest(ctx, rr) // do nothing

	return &order.ReleaseResponse{
		Status: SUCCESS.String(),
	}, nil
}

func processReleaseRequest(ctx context.Context, rr *order.ReleaseRequest) {}

func main() {
	if err := SetupOTel(context.Background()); err != nil {
		log.Fatalf("%v", err)
	}

	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("%v", err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderEngineServer(grpcServer, OrderRPCService{})

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
			semconv.ServiceNameKey.String("order-rpc-service"),
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
