package main

import (
	"context"

	"github.com/google/uuid"
	orderdefs "github.com/sklrsn/gRPC-defs/order"
)

type OrderStatus string

func (os *OrderStatus) String() string {
	return string(*os)
}

var (
	SUCCESS = OrderStatus("SUCCESS")
)

type OrderRPCService struct {
	orderdefs.OrderEngineServer
}

func Reserve(context.Context, *orderdefs.ReserveRequest) (*orderdefs.ReserveResponse, error) {
	return &orderdefs.ReserveResponse{
		OrderReservationId: uuid.NewString(),
	}, nil
}

func Release(context.Context, *orderdefs.ReleaseRequest) (*orderdefs.ReleaseResponse, error) {
	return &orderdefs.ReleaseResponse{
		Status: SUCCESS.String(),
	}, nil
}
