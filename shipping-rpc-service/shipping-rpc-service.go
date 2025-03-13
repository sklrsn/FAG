package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	otellib "github.com/sklrsn/FAG/lib/otel"
	"github.com/sklrsn/gRPC-defs/shipping"
	"go.opentelemetry.io/otel"
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
	otelShutdownFunc, err := otellib.SetupOTel(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer otelShutdownFunc(context.Background())

	listener, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("%v", err)
	}

	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, ShippingRPCService{})

	log.Fatalf("%v", grpcServer.Serve(listener))
}
