package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/sklrsn/gRPC-defs/order"
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

func Reserve(context.Context, *order.ReserveRequest) (*order.ReserveResponse, error) {
	return &order.ReserveResponse{
		OrderReservationId: uuid.NewString(),
	}, nil
}

func Release(context.Context, *order.ReleaseRequest) (*order.ReleaseResponse, error) {
	return &order.ReleaseResponse{
		Status: SUCCESS.String(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("%v", err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderEngineServer(grpcServer, OrderRPCService{})

	log.Fatalf("%v", grpcServer.Serve(listener))
}
