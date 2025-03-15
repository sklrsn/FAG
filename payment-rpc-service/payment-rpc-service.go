package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	otellib "github.com/sklrsn/FAG/lib/otel"
	"github.com/sklrsn/gRPC-defs/payment"
	"go.opentelemetry.io/otel"
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

func (ps PaymentRPCService) Release(ctx context.Context, rr *payment.ReleaseRequest) (*payment.ReleaseResponse, error) {
	ctx, span := tracer.Start(ctx, "Release")
	defer span.End()

	processReleaseRequest(ctx, rr) // do nothing

	return &payment.ReleaseResponse{
		Status: SUCCESS.String(),
	}, nil
}
func processReleaseRequest(ctx context.Context, rr *payment.ReleaseRequest) {} // do nothing

func main() {
	if err := otellib.SetupOTel(context.Background()); err != nil {
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
