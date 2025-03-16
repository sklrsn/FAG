package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/sklrsn/gRPC-defs/payment"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

type PaymentStatus string

func (os *PaymentStatus) String() string {
	return string(*os)
}

var (
	SUCCESS = PaymentStatus("SUCCESS")
)

type PaymentRPCService struct {
	payment.PaymentServer
}

var ServiceName = "payment-rpc-service"

var (
	tracer = otel.Tracer(ServiceName)
)

func (ps PaymentRPCService) PreAuthorize(ctx context.Context, pa *payment.PreAuthorizeRequest) (*payment.PreAuthorizeResponse, error) {
	ctx, span := tracer.Start(ctx, "PreAuthorize")
	defer span.End()

	processPreAuthorizeRequest(ctx, pa) // do nothing

	return &payment.PreAuthorizeResponse{
		PreAuthorizationId: uuid.NewString(),
	}, nil
}

func processPreAuthorizeRequest(ctx context.Context, pa *payment.PreAuthorizeRequest) {} // do nothing

func (ps PaymentRPCService) Capture(ctx context.Context, cr *payment.CaptureRequest) (*payment.CaptureResponse, error) {
	ctx, span := tracer.Start(ctx, "Capture")
	defer span.End()

	processCaptureRequest(ctx, cr) // do nothing

	return &payment.CaptureResponse{
		PaymentCaptureId: uuid.NewString(),
	}, nil
}
func processCaptureRequest(ctx context.Context, pa *payment.CaptureRequest) {} // do nothing

func (ps PaymentRPCService) Release(ctx context.Context, rr *payment.ReimburseRequest) (*payment.ReimburseResponse, error) {
	ctx, span := tracer.Start(ctx, "Release")
	defer span.End()

	processReleaseRequest(ctx, rr) // do nothing

	return &payment.ReimburseResponse{
		Status: SUCCESS.String(),
	}, nil
}
func processReleaseRequest(ctx context.Context, rr *payment.ReimburseRequest) {} // do nothing

func main() {
	if err := SetupOTel(context.Background()); err != nil {
		log.Fatalf("%v", err)
	}

	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("%v", err)
	}

	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServer(grpcServer, PaymentRPCService{})

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
			semconv.ServiceNameKey.String("payment-rpc-service"),
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
