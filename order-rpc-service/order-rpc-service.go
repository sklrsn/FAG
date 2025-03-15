package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	otellib "github.com/sklrsn/FAG/lib/otel"
	"github.com/sklrsn/gRPC-defs/order"
	"go.opentelemetry.io/otel"
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
	if err := otellib.SetupOTel(context.Background()); err != nil {
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
